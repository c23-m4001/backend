//go:build tools

package cmd

import "github.com/spf13/cobra"

const migrationFileContentTemplate = `package migration

func init() {
	sourceDriver.append(
		%s,
		` + "`" + `
		` + "`" + `,
		` + "`" + `
		` + "`" + `,
	)
}
`

func init() {
	rootCmd.AddCommand(migrationFileGen())
}

func migrationFileGen() *cobra.Command {
	var filename string

	cmd := &cobra.Command{
		Use:   "migrate-gen",
		Short: "Generate migration file",
		Long:  "Generate migration file",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.Flags().StringVarP(&filename, "filename", "f", "", "Specify file name without timestamp example: --filename=create_user_table")
	cmd.MarkFlagRequired("filename")

	return cmd
}
