package cmd

import (
	"capstone/config"
	"capstone/delivery/api"
	"capstone/manager"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backend",
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
