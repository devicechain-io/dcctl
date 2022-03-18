/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package cmd

import (
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	v1beta1 "github.com/devicechain-io/dc-k8s/api/v1beta1"
	dck8s "github.com/devicechain-io/dc-k8s/config"
	gen "github.com/devicechain-io/dc-k8s/generators"

	"github.com/pytimer/k8sutil/apply"
	"github.com/spf13/cobra"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"

	"github.com/fatih/color"
)

// Create instance of install command
var installCmd = NewInstallCommand()

// Create command for installing DeviceChain core resources
func NewInstallCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Install DeviceChain components",
		Long:  `Installs DeviceChain k8s manifests and operator`,
		RunE: func(cmd *cobra.Command, args []string) error {
			manifests := dck8s.Manifests()
			crds, err := getManifestContent(manifests)
			if err != nil {
				return err
			}

			fmt.Println("Preparing to install DeviceChain components...")

			dynamicClient, discoveryClient, err := createClients()
			if err != nil {
				return err
			}

			fmt.Println(GreenUnderline("\nInstall Custom Resource Definitions"))
			for _, current := range crds {
				err = applyYaml(dynamicClient, discoveryClient, current.Content)
				if err != nil {
					return err
				}

				fmt.Printf(color.GreenString("Installed resource: %s\n"), color.HiWhiteString(current.Name))
			}

			fmt.Println(GreenUnderline("\nInstall Custom Resources"))
			err = filepath.Walk(GenResFolder, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				b, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				err = applyYaml(dynamicClient, discoveryClient, b)
				if err != nil {
					return err
				}

				fmt.Printf(color.GreenString("Installed resource: %s\n"), color.HiWhiteString(path))
				return nil
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(color.HiGreenString("\nInstallation completed successfully."))
			return nil
		},
	}
}

// Gather all manifest content from the embedded files
func getManifestContent(manifests embed.FS) ([]gen.ConfigurationResource, error) {
	resources := make([]gen.ConfigurationResource, 0)
	err := fs.WalkDir(manifests, "crd/bases", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		f, err := manifests.Open(path)
		if err != nil {
			return err
		}
		info, err := f.Stat()
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		b, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		resources = append(resources, gen.ConfigurationResource{
			Name:    path,
			Content: b,
		})
		return nil
	})
	return resources, err
}

// Create k8s clients needed to apply resources
func createClients() (dynamic.Interface, *discovery.DiscoveryClient, error) {
	dynamicClient, err := dynamic.NewForConfig(v1beta1.ClientConfig)
	if err != nil {
		return nil, nil, err
	}
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(v1beta1.ClientConfig)
	if err != nil {
		return nil, nil, err
	}

	// You can add other(crd/build-in) resource scheme
	// utilruntime.Must(imagepolicyv1alpha1.AddToScheme(apply.Scheme))
	return dynamicClient, discoveryClient, nil
}

// Apply yaml to k8s
func applyYaml(dynamicClient dynamic.Interface, discoveryClient *discovery.DiscoveryClient, yaml []byte) error {
	applyOptions := apply.NewApplyOptions(dynamicClient, discoveryClient)
	if err := applyOptions.Apply(context.TODO(), []byte(yaml)); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(installCmd)
}
