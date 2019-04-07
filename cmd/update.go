package cmd

import (
	"github.com/alexander-demichev/ocp-mig-test-data-cli/repo"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update ocp-mig-test-data command",
	Run: func(cmd *cobra.Command, args []string) {
		repo.UpdateRepo()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
