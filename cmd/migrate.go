//go:build tools

package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(migrateCommand())
}

func migrateCommand() *cobra.Command {
	var isRollingBack bool
	var steps int

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database table",
		Long:  "Migrate database table using steps or rollback",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.Flags().BoolVarP(&isRollingBack, "rollback", "r", false, "Rollback migration")
	cmd.Flags().IntVarP(&steps, "steps", "s", 0, "Specify steps for 'n' migration (default: 0)")

	return cmd
}
