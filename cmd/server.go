package cmd

import (
	"github.com/spf13/cobra"
)

func newServerCommand() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "server operations",
	}

	serverCmd.PersistentFlags().StringP("config", "c", "", "Path to config file")
	serverCmd.AddCommand(newMigrationCmd())
	serverCmd.AddCommand(newStartCmd())
	return serverCmd
}
