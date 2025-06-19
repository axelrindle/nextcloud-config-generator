package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nc-cfg-gen",
	Short: "Generate configuration files for Nextcloud.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	rootCmd.AddGroup(
		&cobra.Group{ID: "generate", Title: "Config Generation"},
	)
}
