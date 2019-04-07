package cmd

import (
	"errors"
	"log"

	"github.com/alexander-demichev/ocp-mig-test-data-cli/ansible"

	"github.com/alexander-demichev/ocp-mig-test-data-cli/list"
	"github.com/spf13/cobra"
)

var cfgFile string
var all bool
var withBackup bool
var withRestore bool
var deploy bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ocp-mig-test-data-cli",
	Short: "CLI for deploying ocp-mig-test-data",
	Long: `CLI for deploying ocp-mig-test-data with ability to create backups and restores.
	Use ocp-mig-test-data-cli help to get more info`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		withBackup, _ := cmd.Flags().GetBool("restore")
		withRestore, _ := cmd.Flags().GetBool("backup")
		withData, _ := cmd.Flags().GetBool("deploy")

		if all {
			ansible.RunPlaybook(args[0], true, true, true)
			return
		}

		if !withData && !withBackup && !withRestore {
			ansible.RunPlaybook(args[0], true, false, false)
		}
		if len(args) > 1 {
			ansible.RunPlaybook(args[0], withData, withBackup, withRestore)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("Not possible to deploy more then 1 sample")
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
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "deploy sample, create backup and restore")
	rootCmd.Flags().BoolVarP(&withBackup, "backup", "b", false, "create backup")
	rootCmd.Flags().BoolVarP(&withRestore, "restore", "r", false, "create restore")
	rootCmd.Flags().BoolVarP(&deploy, "deploy", "d", false, "create restore")
}
