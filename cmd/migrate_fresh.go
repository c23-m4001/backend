//go:build tools

package cmd

import (
	"capstone/config"
	"capstone/manager"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateFreshCommand())
}

func migrateFreshCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-fresh",
		Short: "Refresh all database tables",
		Long:  "Refresh all database tables",
		Run: func(cmd *cobra.Command, args []string) {
			config.DisableDebug()

			container := manager.NewContainer(manager.LoadDefault)

			if err := container.InfrastructureManager().RefreshDB(); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
