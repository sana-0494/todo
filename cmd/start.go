package cmd

import (
	"context"
	"fmt"
	"todo/server"

	"github.com/spf13/cobra"
)

func newStartCmd() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start server",
		RunE:  start,
	}
	return startCmd
}

func start(cmd *cobra.Command, _ []string) error {
	ctx := context.Background()
	cfgPath, err := cmd.Flags().GetString("config")
	if err != nil {
		fmt.Println("invlaid flag for migrate command")
		return err
	}
	
	cfg := getConfig(cfgPath)
	server.Serve(ctx, cfg)
	return nil
}
