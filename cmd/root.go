package cmd

import (
	"errors"

	"github.com/alexander-demichev/ocp-mig-test-data-cli/pkg/ansible"
	"github.com/alexander-demichev/ocp-mig-test-data-cli/pkg/log"

	"github.com/alexander-demichev/ocp-mig-test-data-cli/pkg/list"
	"github.com/spf13/cobra"
)

var cfgFile string
var all bool
var withBackup bool
var withRestore bool
var withSample bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ocp-mig-test-data-cli",
	Short: "CLI for creating ocp-mig-test-data",
	Long: `CLI for creating ocp-mig-test-data with ability to create backups and restores.
	Use ocp-mig-test-data-cli help to get more info`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		withBackup, _ := cmd.Flags().GetBool("backup")
		withRestore, _ := cmd.Flags().GetBool("restore")
		withSample, _ := cmd.Flags().GetBool("sample")

		if all {
			ansible.RunPlaybook(args[0], true, true, true)
			return
		}
		if !withSample && !withBackup && !withRestore {
			ansible.RunPlaybook(args[0], true, false, false)
		}
		if len(args) == 1 {
			ansible.RunPlaybook(args[0], withSample, withBackup, withRestore)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("Not possible to create more then 1 sample")
		}

		if !list.IsValidPlaybook(args[0]) {
			return errors.New("This sample doesn't exists")
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "create sample, backup and restore")
	rootCmd.Flags().BoolVarP(&withBackup, "backup", "b", false, "create backup")
	rootCmd.Flags().BoolVarP(&withRestore, "restore", "r", false, "create restore")
	rootCmd.Flags().BoolVarP(&withSample, "sample", "s", false, "create sample")
}
