/*
Copyright © 2022 DANIEL DE JAGER <daniel.dejager.au@gmail.com>

*/
package configure

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const maskLimit = 4

// ConfigureCmd represents the configure command
var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "The configure command allows you to configure third party applications that the CLI interfaces with",
	Long:  `The SECOPS CLI is a fully functional Command Line Interface (CLI) that interacts with a variety of security related applications`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SetConfigProperty(propName, propValue string) {
	viper.Set(propName, propValue)
	if viperErr := viper.SafeWriteConfig(); viperErr != nil {
		_ = viper.WriteConfig()
	}
}

func MaskKey(key string) string {
	if len(key) > maskLimit {
		return "******" + key[len(key)-4:]
	} else if len(key) > 1 {
		return "******"
	} else {
		return ""
	}
}
