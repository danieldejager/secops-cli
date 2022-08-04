/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package configure

import (
	"fmt"

	"github.com/danieldejager/secops-cli/internal/params"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiKey string
)

// virustotalCmd represents the virustotal command
var virustotalCmd = &cobra.Command{
	Use:   "virustotal",
	Short: "The virustotal command allows you to configure keys required to work with the VirusTotal API",
	Long:  `The SECOPS CLI is a fully functional Command Line Interface (CLI) that interacts with a variety of security related applications`,
	Run: func(cmd *cobra.Command, args []string) {
		SetConfigProperty(params.VirusTotalAPIKey, apiKey)
		fmt.Printf("Configured API Key [%s]", MaskKey(viper.GetString(params.VirusTotalAPIKey)))
	},
}

func init() {

	virustotalCmd.Flags().StringVarP(&apiKey, "apikey", "k", "", "Sets the value of the API Key to use with VirusTotal")

	virustotalCmd.MarkFlagRequired("apikey")

	ConfigureCmd.AddCommand(virustotalCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// virustotalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// virustotalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
