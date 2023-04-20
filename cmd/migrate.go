// //go:build tools

package cmd

import (
	"capstone/config"
	"capstone/manager"

	"github.com/spf13/cobra"
)

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
			config.DisableDebug()

			container := manager.NewContainer(manager.LoadDefault)
			if err := container.InfrastructureManager().MigrateDB(isRollingBack, steps); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&isRollingBack, "rollback", "r", false, "Rollback migration")
	cmd.Flags().IntVarP(&steps, "steps", "s", 0, "Specify steps for 'n' migration (default: 0)")

	return cmd
}
