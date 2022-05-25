/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// Create common command for creating DeviceChain resources
var constructionDataCmd = &cobra.Command{
	Use:   "construction",
	Short: "Bootstrap construction sample data",
	Long:  `Bootstraps system microservices with sample data for construction use case`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return boostrapConstructionData()
	},
}

// Bootstraps system microservices with construction sample dataset.
func boostrapConstructionData() error {
	return nil
}

func init() {
	bootstrapCmd.AddCommand(constructionDataCmd)
}
