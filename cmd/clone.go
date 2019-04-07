package cmd

import (
	"github.com/alexander-demichev/mig-data-cli/pkg/repo"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone ocp-mig-test-data repo",
	Run: func(cmd *cobra.Command, args []string) {
		repo.CloneRepo()
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
