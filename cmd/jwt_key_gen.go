//go:build tools

package cmd

import (
	"capstone/config"
	"capstone/util"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(jwtKeyGenCommand())
}

func jwtKeyGenCommand() *cobra.Command {
	var isForce bool

	cmd := &cobra.Command{
		Use:   "jwt-key-gen",
		Short: "Generate rsa key pair for jwt",
		Long:  "This command is used to generat rsa key pair for jwt",
		Run: func(cmd *cobra.Command, args []string) {
			isFileExists := util.IsFileExist(config.GetJwtPrivateKeyFilePath())

			folderPath := filepath.Dir(config.GetJwtPrivateKeyFilePath())
			os.MkdirAll(folderPath, os.ModePerm)

			if isForce || !isFileExists {
				bitSize := 4096

				// Generate RSA key
				key, err := rsa.GenerateKey(rand.Reader, bitSize)
				if err != nil {
					panic(err)
				}

				// Extract public component
				pub := key.Public()

				// Encode private key using PEM
				keyPEM := pem.EncodeToMemory(
					&pem.Block{
						Type:  "RSA PRIVATE KEY",
						Bytes: x509.MarshalPKCS1PrivateKey(key),
					},
				)

				// Encode public key using PEM
				pubPEM := pem.EncodeToMemory(&pem.Block{
					Type:  "RSA PUBLIC KEY",
					Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
				})

				// write to file
				if err := ioutil.WriteFile(config.GetJwtPrivateKeyFilePath(), keyPEM, 0700); err != nil {
					panic(err)
				}
				if err := ioutil.WriteFile(config.GetJwtPublicKeyFilePath(), pubPEM, 0755); err != nil {
					panic(err)
				}
				if err := ioutil.WriteFile(config.GetJwtGitIgnoreFilePath(), []byte("*\n!.gitignore\n"), 0755); err != nil {
					panic(err)
				}

			}

		},
	}

	cmd.Flags().BoolVarP(&isForce, "force", "f", false, "force recreate rsa key pair for jwt")

	return cmd
}
