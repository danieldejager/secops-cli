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
var ConfigureVirustotalCmd = &cobra.Command{
	Use:   "virustotal",
	Short: "The virustotal command allows you to configure keys required to work with the VirusTotal API",
	Long:  `The SECOPS CLI is a fully functional Command Line Interface (CLI) that interacts with a variety of security related applications`,
	Run: func(cmd *cobra.Command, args []string) {
		SetConfigProperty(params.VirusTotalAPIKey, apiKey)
		fmt.Printf("Configured VirusTotal API Key [%s]", MaskKey(viper.GetString(params.VirusTotalAPIKey)))
	},
}

func init() {

	ConfigureVirustotalCmd.Flags().StringVarP(&apiKey, "apikey", "k", "", "Sets the value of the API Key to use with VirusTotal")

	ConfigureVirustotalCmd.MarkFlagRequired("apikey")

	ConfigureCmd.AddCommand(ConfigureVirustotalCmd)

}
