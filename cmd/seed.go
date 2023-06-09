//go:build tools

package cmd

import (
	"capstone/config"
	"capstone/database/seeder"
	"capstone/database/seeder/production_seeder"
	"capstone/manager"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCommand())
}

func seedCommand() *cobra.Command {
	var flagProduction bool

	cmd := &cobra.Command{
		Use:   "seed",
		Short: "Seed all database",
		Long:  "Seed all database with either production or development data",
		Run: func(cmd *cobra.Command, args []string) {
			config.DisableDebug()

			container := manager.NewContainer(manager.LoadDefault)
			defer func() {
				if err := container.Close(); err != nil {
					panic(err)
				}
			}()

			repositoryManager := container.RepositoryManager()

			if len(args) > 0 {
				fmt.Printf("Seeder for table `%s` not found\n", args[0])
				return
			}

			if flagProduction {
				production_seeder.SeedAll(repositoryManager)
			} else {
				seeder.SeedAll(repositoryManager)
			}
		},
	}

	cmd.Flags().BoolVarP(&flagProduction, "production", "p", false, "seed data using production data")

	// loop add command seed per table

	return cmd
}

func seedTableCommand(tableName string) *cobra.Command {
	var flagProduction bool

	cmd := &cobra.Command{
		Use:   tableName,
		Short: fmt.Sprintf("Seed %s table", tableName),
		Long:  fmt.Sprintf("Seed %s table with either production or development data", tableName),
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	// TODO: need to check exists production_seeder exists
	cmd.Flags().BoolVarP(&flagProduction, "production", "p", false, fmt.Sprintf("seed %s table using production data", tableName))

	return cmd
}
