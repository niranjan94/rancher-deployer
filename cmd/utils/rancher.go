package utils

import (
	"strings"
	"encoding/json"
	"fmt"
)

type RancherServerConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	TokenKey  string `json:"tokenKey"`
	Url       string `json:"url"`
	Project   string `json:"project"`
	Cacert    string `json:"cacert"`
}

type RancherServersConfig struct {
	PrimaryServer RancherServerConfig
}

type RancherConfig struct {
	Servers RancherServersConfig
	CurrentServer string
}

func GetRancherConfig(token string, project string, serverUrl string) string {

	splitToken := strings.Split(token, ":")

	config := RancherConfig{
		CurrentServer: "PrimaryServer",
		Servers: RancherServersConfig{
			PrimaryServer: RancherServerConfig{
				AccessKey: splitToken[0],
				SecretKey: splitToken[1],
				TokenKey:  token,
				Url:       serverUrl,
				Project:   project,
				Cacert:    "",
			},
		},
	}

	jsonByteArray, err := json.Marshal(config)

	if err != nil {
		panic(fmt.Errorf("Could not marshal as json: %s \n", err))
		return ""
	}

	return string(jsonByteArray[:])
}