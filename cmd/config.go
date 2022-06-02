/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	key, value string
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Change your config settings",
	Long:  ``,
}

// configListCmd represents the config command
var configListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your config settings",
	Long:    ``,
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		user := []string{"email", "name"}
		printConfigs("user", user)
	},
}

// configSetCmd represents the config command
var configSetCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set your config settings",
	Long:    ``,
	Example: "  lbctl config set user.name Michele Lefebvre",
	Run: func(cmd *cobra.Command, args []string) {
		var configKey, configValue string
		if len(args) == 0 && key == "" {
			fmt.Println("you must enter the config key you would like to modify")
			os.Exit(1)
		} else if key != "" {
			configKey = key
		} else {
			configKey = args[0]
		}

		if len(args) <= 1 && value == "" {
			fmt.Println("you must enter the config value to modify")
			os.Exit(1)
		} else if value != "" {
			configValue = value
		} else {
			for i, s := range args {
				if i == 0 {
					continue
				} else if i == 1 {
					configValue = s
				} else {
					configValue = configValue + " " + s
				}
			}
		}

		viper.Set(configKey, configValue)
		viper.WriteConfig()
		fmt.Println(configKey, "changed to", configValue)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configSetCmd)

	configSetCmd.Flags().StringVarP(&key, "key", "k", "", "The name of the config you want to set")
	configSetCmd.Flags().StringVarP(&value, "value", "v", "", "The new value to asign the config")
}

func printConfigs(configType string, configs []string) {
	for _, c := range configs {
		s := fmt.Sprintf("%v.%v", configType, c)
		fmt.Printf("%v=%v\n", s, viper.Get(s))
	}
}
