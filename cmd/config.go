package cmd

import (
	"github.com/spf13/viper"
	"fmt"
	"strings"
)

func LoadConfig() {
	replacer := strings.NewReplacer(".", "_")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("DEPLOYER")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(".rancher.deployer")
	viper.AddConfigPath("$HOME/.rancher-deployer")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}