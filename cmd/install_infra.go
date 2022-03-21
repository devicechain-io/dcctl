/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	corev1beta1 "github.com/devicechain-io/dc-k8s/api/v1beta1"
	"github.com/fatih/color"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

const (
	NS_DC_SYSTEM = "dc-system"
)

// Create instance of install infra command
var installInfraCmd = NewInstallInfraCommand()

// Create command for installing DeviceChain infrastructure
func NewInstallInfraCommand() *cobra.Command {
	return &cobra.Command{
		Use:          "infra",
		Short:        "Install DeviceChain infrastructure components",
		Long:         `Installs and configures DeviceChain infrastructure dependencies`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return installInfraComponents()
		},
	}
}

// Install all infrastructure components
func installInfraComponents() error {
	// Validate that system namespace exists.
	err := assureSystemNamespace()
	if err != nil {
		return err
	}

	// Locate and/or setup Helm repositories.
	settings := cli.New()
	rfile, err := assureHelmRepositoryConfig(settings)
	if err != nil && !os.IsExist(err) {
		return err
	}
	// Test adding an entry.
	entry := &repo.Entry{
		Name: "bitnami",
		URL:  "https://charts.bitnami.com/bitnami",
	}

	err = addHelmRepository(entry, settings, rfile)
	if err != nil {
		return err
	}

	return nil
}

// Assure that
func assureSystemNamespace() error {
	// Check for existing namespace.
	fmt.Print(color.WhiteString("Verifying DeviceChain system namespace... "))
	ns := &corev1.Namespace{}
	err := corev1beta1.V1Client.Get(context.Background(), types.NamespacedName{Name: NS_DC_SYSTEM}, ns)
	if err != nil {
		// Attempt to create the namespace.
		ns = &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: NS_DC_SYSTEM}}
		err = corev1beta1.V1Client.Create(context.Background(), ns)
		fmt.Println(color.HiGreenString("Created system namespace."))
	} else {
		fmt.Println(color.HiGreenString("System namespace verified."))
	}
	return err
}

// Assure that Helm repository is initialized.
func assureHelmRepositoryConfig(settings *cli.EnvSettings) (*repo.File, error) {
	// Make sure Helm repository path exists
	err := os.MkdirAll(filepath.Dir(settings.RepositoryConfig), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	// Get or create the repository file
	var file *repo.File
	if _, err := os.Stat(settings.RepositoryConfig); errors.Is(err, os.ErrNotExist) {
		file = repo.NewFile()
		err = file.WriteFile(settings.RepositoryConfig, 0755)
		if err != nil {
			return nil, err
		}
		fmt.Println(color.WhiteString("Created new Helm repositories configuration at: "),
			color.HiGreenString(settings.RepositoryConfig))
	} else {
		file, err = repo.LoadFile(settings.RepositoryConfig)
		if err != nil {
			return nil, err
		}
		fmt.Println(color.WhiteString("Using existing Helm repositories configuration at: "),
			color.HiGreenString(settings.RepositoryConfig))
	}
	return file, nil
}

// Add Helm repository to configuration
func addHelmRepository(entry *repo.Entry, settings *cli.EnvSettings, rfile *repo.File) error {
	fmt.Printf(color.WhiteString("Checking repository '%s' ... "), entry.Name)
	if rfile.Has(entry.Name) {
		fmt.Println(color.HiGreenString("FOUND"))
		return nil
	}

	// Pull index file to verify..
	r, err := repo.NewChartRepository(entry, getter.All(settings))
	if err != nil {
		return err
	}
	_, err = r.DownloadIndexFile()
	if err != nil {
		return err
	}

	// Update repositories list.
	rfile.Update(entry)
	err = rfile.WriteFile(settings.RepositoryConfig, 0755)
	if err != nil {
		return err
	}
	fmt.Println(color.HiGreenString("ADDED"))

	return nil
}

func init() {
	rootCmd.AddCommand(installInfraCmd)
}
