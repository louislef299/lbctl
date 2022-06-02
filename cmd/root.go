/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>

*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "lbctl",
	Short: "Little Bitta Control",
	Long: `This cli tool will be used to control the little bitta
cloud environments. As it is built out, its use will
become more clear.

                        ____
                   .---'-    \
      .-----------/           \
     /           (         ^  |   __
&   (             \        O  /  / .'
'._/(              '-'  (.   (_.' /
     \                    \     ./
      |    |       |    |/ '._.'
       )   @).____\|  @ |
   .  /    /       (    | 
  \|, '_:::\  . ..  '_:::\ ..\).`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := path.Join(home, ".lbctl")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", configPath, "Path to config file")

	configInit()
}

func configInit() {
	configName := "config"
	viper.SetConfigName(configName)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath(configFile)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found, check if config folder exists
			if _, err := os.Stat(configFile); os.IsNotExist(err) {
				err = os.Mkdir(configFile, 0744)
				if err != nil {
					log.Fatal(err)
				}
			}
			interactive(configFile, configName)
		} else {
			log.Fatal(err)
		}
	}
}

func interactive(configPath, configName string) {
	fmt.Printf("enter your full name(first last): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("An error occured while reading input. Please try again", err)
	}
	input = strings.TrimSuffix(input, "\n")
	viper.Set("user.name", input)
	fmt.Printf("enter your email: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal("An error occured while reading input. Please try again", err)
	}
	input = strings.TrimSuffix(input, "\n")
	viper.Set("user.email", input)
	viper.WriteConfigAs(path.Join(configPath, configName))
}
