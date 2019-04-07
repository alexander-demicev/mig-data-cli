package cmd

import (
	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "A help command",
	Run: func(cmd *cobra.Command, args []string) {
		helpMessage()
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}

func helpMessage() {
	println(` It's required to run 'test-data-cli clone before deploying samples'
Usage: 'test-data-cli {sample-name} {actions}'
Examples: 'test-data-cli mysql-pvc backup restore', 'test-data-cli mysql-pvc all'
clone - Clone ocp-mig-test data repo
update - update ocp-mit-test data repo
list - List all samples
backup - Create velero backup
restore - Create velero restore
all - Deploy sample, create backup and restore
help - Help message
version - Get version info
`)
}
