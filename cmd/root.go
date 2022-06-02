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

	lberror "github.com/louislef299/lbctl/error"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	configInit()
}

func configInit() {
	home, err := homedir.Dir()
	lberror.CheckError(err)
	configPath := path.Join(home, ".lbctl")

	configName := "config"
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found, check if config folder exists
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				err = os.Mkdir(configPath, 0744)
				if err != nil {
					log.Fatal(err)
				}
			}
			interactive(configPath, configName)
		} else {
			log.Fatal(err)
		}
	}
	name := viper.Get("name")
	email := viper.Get("email")
	if name == nil || email == nil {
		interactive(configPath, configName)
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
	viper.Set("name", input)
	fmt.Printf("enter your email: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal("An error occured while reading input. Please try again", err)
	}
	input = strings.TrimSuffix(input, "\n")
	viper.Set("email", input)
	viper.WriteConfigAs(path.Join(configPath, configName))
}
