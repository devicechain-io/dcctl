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
	"strings"

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
var installCoreCmd = NewInstallCoreCommand()

// Create command for installing DeviceChain core resources
func NewInstallCoreCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "core",
		Short: "Install core components",
		Long:  `Installs Kubernetes manifests and operator`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Preparing to install DeviceChain core components...")

			dynamicClient, discoveryClient, err := createClients()
			if err != nil {
				return err
			}

			// Install CRDs.
			err = installCrds(dynamicClient, discoveryClient)
			if err != nil {
				return err
			}

			// Install RBAC files.
			err = installRbac(dynamicClient, discoveryClient)
			if err != nil {
				return err
			}

			// Install operator files.
			err = installOperator(dynamicClient, discoveryClient)
			if err != nil {
				return err
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

				fmt.Printf(color.WhiteString("Installed resource: %s\n"), color.GreenString(path))
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

// Install all custom resource definitions from k8s metadata.
func installCrds(dynamicClient dynamic.Interface, discoveryClient *discovery.DiscoveryClient) error {
	fmt.Println(GreenUnderline("\nInstall Custom Resource Definitions"))
	crdfiles := dck8s.CrdFiles()
	crds, err := getEmbeddedContent(crdfiles, "crd/bases")
	if err != nil {
		return err
	}
	for _, current := range crds {
		err = applyYaml(dynamicClient, discoveryClient, current.Content)
		if err != nil {
			return err
		}

		fmt.Printf(color.WhiteString("Installed CRD: %s\n"),
			color.GreenString(strings.TrimPrefix(current.Name, "crd/bases/")))
	}
	return nil
}

// Install all RBAC definitions from k8s metadata.
func installRbac(dynamicClient dynamic.Interface, discoveryClient *discovery.DiscoveryClient) error {
	fmt.Println(GreenUnderline("\nInstall RBAC Components"))
	crdfiles := dck8s.RbacFiles()
	crds, err := getEmbeddedContent(crdfiles, "rbac")
	if err != nil {
		return err
	}
	for _, current := range crds {
		err = applyYaml(dynamicClient, discoveryClient, current.Content)
		if err != nil {
			return err
		}

		fmt.Printf(color.WhiteString("Installed RBAC: %s\n"),
			color.GreenString(strings.TrimPrefix(current.Name, "rbac/")))
	}
	return nil
}

// Install all operator definitions from k8s metadata.
func installOperator(dynamicClient dynamic.Interface, discoveryClient *discovery.DiscoveryClient) error {
	fmt.Println(GreenUnderline("\nInstall Operator Components"))
	mgrfiles := dck8s.ManagerFiles()
	mgrs, err := getEmbeddedContent(mgrfiles, "manager")
	if err != nil {
		return err
	}
	for _, current := range mgrs {
		err = applyYaml(dynamicClient, discoveryClient, current.Content)
		if err != nil {
			return err
		}

		fmt.Printf(color.WhiteString("Installed Operator Component: %s\n"),
			color.GreenString(strings.TrimPrefix(current.Name, "manager/")))
	}
	return nil
}

// Gather all content from the embedded files in the relative path.
func getEmbeddedContent(embedded embed.FS, path string) ([]gen.ConfigurationResource, error) {
	resources := make([]gen.ConfigurationResource, 0)
	err := fs.WalkDir(embedded, path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		f, err := embedded.Open(path)
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
		if strings.HasPrefix(d.Name(), "kust") {
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
	installCmd.AddCommand(installCoreCmd)
}
