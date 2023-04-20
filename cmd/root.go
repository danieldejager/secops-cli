/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/MakeNowJust/heredoc"
	"github.com/danieldejager/secops-cli/internal/commands"
	"github.com/danieldejager/secops-cli/internal/commands/configure"
	"github.com/danieldejager/secops-cli/internal/params"
	"github.com/danieldejager/secops-cli/internal/wrappers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const cfgType = "yaml"
const cfgFile = ".secops-cli"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func NewSecOpsCLI(
	virusTotalWrapper wrappers.VirusTotalWrapper,
) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "secops-cli",
		Short: "The secops-cli is a tool that automates tasks for security analysts before, during and after a security incident",
		Long:  `The SECOPS CLI is a fully functional Command Line Interface (CLI) that interacts with a variety of security related applications`,
		Example: heredoc.Doc (
			`
			$secops-cli configure virustotal --apikey my_api_key
			$secops-cli virustotal --filehash my_file_hash 
			`,
		),
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

 

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".secops-cli" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType(cfgType)
	viper.SetConfigName(cfgFile)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config File Updated")
	}
}
