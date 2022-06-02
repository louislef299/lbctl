/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package cmd

import (
	"github.com/louislef299/lbctl/lberror"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize lbctl",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		lberror.CheckError(err)
		/*
			- check if config folder exists
			- if not, create it with config file
			- request for name, email
		*/
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
