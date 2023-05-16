//go:build tools

package cmd

import (
	"capstone/util"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

const migrationFilePath string = "database/migration"

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
			var (
				version           = util.CurrentDateTime().Format("20060102150405")
				migrationFilePath = migrationFilePath + "/" + fmt.Sprintf("%s_%s.go", version, filename)
			)

			if err := ioutil.WriteFile(migrationFilePath, []byte(fmt.Sprintf(migrationFileContentTemplate, version)), 0644); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&filename, "filename", "f", "", "Specify file name without timestamp example: --filename=create_user_table")
	cmd.MarkFlagRequired("filename")

	return cmd
}
