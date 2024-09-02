package cmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "todo <command> <arguments> [flags]",
		Short: "create, delete, update your todo list",
	}
	cmd.AddCommand(newServerCommand())
	return cmd
}
