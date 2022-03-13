/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	dm "github.com/devicechain-io/dc-devicemanagement/config"
	dci "github.com/devicechain-io/dc-microservice/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const GenResFolder = "resources"

// Create common command for creating DeviceChain resources
var resgenCmd = &cobra.Command{
	Use:   "resgen",
	Short: "Generate configuration resources",
	Long:  `Generates configuration resources directly from the microservice codebase`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating resources from source code...")
		os.MkdirAll(GenResFolder, 0777)

		// Generate instance resources
		fmt.Println(GreenUnderline("\nInstance Resources"))
		err := generateInstanceResources()
		if err != nil {
			return err
		}

		// Generate resources for each microservice
		fmt.Println(GreenUnderline("\nMicroservice Resources"))
		err = generateMicroserviceResources(dm.ResourceProvider{})
		if err != nil {
			return err
		}
		return nil
	},
}

// Generate instance resources by introspecting microservice configs
func generateInstanceResources() error {
	dcires, err := dci.GetConfigurationResources()
	if err != nil {
		return err
	}
	for _, dci := range dcires {
		path := filepath.Join(GenResFolder, fmt.Sprintf("%s.yaml", dci.Name))
		err = os.WriteFile(path, dci.Content, 0644)
		if err != nil {
			return err
		}
		fmt.Printf(color.GreenString("Generated instance resource: %s\n"), color.HiWhiteString(path))
	}
	fmt.Println()
	return nil
}

// Generate microservice resources by introspecting microservice configs
func generateMicroserviceResources(prov dci.ConfigurationResourceProvider) error {
	dcires, err := prov.GetConfigurationResources()
	if err != nil {
		return err
	}
	for _, dci := range dcires {
		path := filepath.Join(GenResFolder, fmt.Sprintf("%s.yaml", dci.Name))
		err = os.WriteFile(path, dci.Content, 0644)
		if err != nil {
			return err
		}
		fmt.Printf(color.GreenString("Generated microservice resource: %s\n"), color.HiWhiteString(path))
	}
	fmt.Println()
	return nil
}

func init() {
	rootCmd.AddCommand(resgenCmd)
}
