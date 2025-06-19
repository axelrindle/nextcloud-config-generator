package cli

import (
	"log"
	"time"

	"github.com/axelrindle/nc-cfg-gen/builder/docker"
	"github.com/axelrindle/nc-cfg-gen/nextcloud"
	"github.com/axelrindle/nc-cfg-gen/print"
	"github.com/axelrindle/nc-cfg-gen/types"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var driver string

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Generates the secret values instanceid, passwordsalt and secret.",
	Long: `Nextcloud operates with three secret parameters: instanceid, passwordsalt and secret.
			These are required for Nextcloud to run and only generated once. The values must never change,
			otherwise all data - especially when encrypted - will be rendered unusable.`,
	GroupID: "generate",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[37], 100*time.Millisecond)
		s.Suffix = " Generating NextCloud configuration …"
		s.Start()

		var cfg *nextcloud.ConfigSecrets
		var err *types.Error

		if driver == "docker" {
			// log.Println("Using Docker for setup …")
			cfg, err = docker.GenerateConfig()
		} else if driver == "kubernetes" {
			// log.Println("Using Kubernetes for setup …")
		} else {
			log.Fatalf("Invalid driver %s", driver)
		}

		s.Stop()

		if err != nil {
			err.Throw()
		}

		printer := print.Printer{}
		printer.PrintHead()
		printer.PrintString(cfg.InstanceID, "instanceid")
		printer.PrintString(cfg.Secret, "secret")
		printer.PrintString(cfg.PasswordSalt, "passwordsalt")
		printer.PrintFoot()
	},
}

func init() {
	secretsCmd.Flags().StringVar(&driver, "driver", "", "The driver to use (available: docker, kubernetes)")

	secretsCmd.MarkFlagRequired("driver")

	rootCmd.AddCommand(secretsCmd)
}
