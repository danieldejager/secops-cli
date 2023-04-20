/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"os"

	"github.com/danieldejager/secops-cli/cmd"
	"github.com/danieldejager/secops-cli/internal/params"
	"github.com/danieldejager/secops-cli/internal/wrappers"
	"github.com/danieldejager/secops-cli/wrappers/configuration"
	"github.com/spf13/viper"
)

func main() {

	// load the keys and env
	bindKeysToEnvAndDefault()
	configuration.LoadConfiguration()

	virustotal := viper.GetString(params.VTBasePathKey)

	virustotalWrapper = wrappers.NewHTTPVirusTotalWrapper(virustotal)

	cmd.Execute()
}

func bindKeysToEnvAndDefault() {
	for _, b := range params.EnvVarsBinds {
		err := viper.BindEnv(b.Key, b.Env)
		if err != nil {
			exitIfError(err)
		}
	}
}

func exitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
