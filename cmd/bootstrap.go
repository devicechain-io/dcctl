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
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap system data",
	Long:  `Bootstraps system microservices with example datasets`,
}

func init() {
	bootstrapCmd.PersistentFlags().StringP("server", "s", "localhost", "server hostname targeted for remote calls")
	bootstrapCmd.PersistentFlags().StringP("instance", "i", "dc1", "instance id targeted for remote calls")
	bootstrapCmd.PersistentFlags().StringP("tenant", "t", "tenant1", "tenant id targeted for remote calls")

	rootCmd.AddCommand(bootstrapCmd)
}
