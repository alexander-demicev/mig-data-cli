package cmd

import (
	"github.com/alexander-demichev/ocp-mig-test-data-cli/pkg/list"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		playbookList := list.List()
		list.PrintList(playbookList)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
