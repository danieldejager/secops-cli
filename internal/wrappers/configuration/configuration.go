package configuration

import (
	"log"
	"os/user"

	"github.com/spf13/viper"
)

const configFile = "secopscli"
const configType = "yaml"

func LoadConfiguration() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Cannot find user")
	}
	
	viper.SetConfigName(configFile)
	viper.SetConfigType(configType)
	viper.AddConfigPath(usr.HomeDir)
	errs := viper.ReadInConfig()
	if errs != nil {
		log.Fatal("Cannot find config file")
	}
}