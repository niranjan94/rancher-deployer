package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

func LoadConfig(pathToConfig string) {
	replacer := strings.NewReplacer(".", "_")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("DEPLOYER")
	viper.SetEnvKeyReplacer(replacer)

	if pathToConfig != "" {
		viper.SetConfigType("yaml")
		fileContent, err := ioutil.ReadFile(pathToConfig)
		if err != nil {
			panic(fmt.Errorf("Failed to read config file at %s\nError: %s", fileContent, err))
		}
		err = viper.ReadConfig(bytes.NewBuffer(fileContent))
		if err != nil {
			panic(fmt.Errorf("Failed to read config file at %s\nError: %s", fileContent, err))
		}
		return
	}

	viper.SetConfigName(".rancher.deployer")
	viper.AddConfigPath("$HOME/.rancher-deployer")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config file: %s \n", err))
	}
}