package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "Run server",
	Long:  `Run backend server`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
