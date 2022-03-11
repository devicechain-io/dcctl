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

	"github.com/devicechain-io/dc-microservice/config"
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
		os.MkdirAll(GenResFolder, 0777)
		return generateInstanceResources()
	},
}

// Generate instance resources by introspecting microservice configs
func generateInstanceResources() error {
	dcires, err := config.GetInstanceConfigurationResources()
	if err != nil {
		return err
	}
	fmt.Println("Generating resources from source code...")
	fmt.Println(GreenUnderline("\nInstance Resources"))
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

func init() {
	rootCmd.AddCommand(resgenCmd)
}
