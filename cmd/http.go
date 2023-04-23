//go:build http

package cmd

import (
	"capstone/config"
	"capstone/delivery/api"
	"capstone/manager"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(routerCommand())
}

func routerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Run server",
		Long:  `Run backend server`,
		Run: func(cmd *cobra.Command, args []string) {
			container := manager.NewContainer(manager.LoadDefault)
			defer func() {
				if err := container.Close(); err != nil {
					panic(err)
				}
			}()

			router := api.NewRouter(container)

			addr := fmt.Sprintf(":%d", config.GetConfig().Port)
			err := router.Run(addr)
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd
}
