//go:build tools

package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(migrateFreshCommand())
}

func migrateFreshCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-fresh",
		Short: "Refresh all database tables",
		Long:  "Refresh all database tables",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
