package rancher

import (
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"os"
)

type serverConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	TokenKey  string `json:"tokenKey"`
	Project   string `json:"project"`
	Cacert    string `json:"cacert"`
	Url       string `json:"url"`
}

type serversConfig struct {
	PrimaryServer serverConfig
}

type config struct {
	CurrentServer string
	Servers       serversConfig
}

type Rancher struct {
	homeDirectory string
}

func New(homeDirectory string, token string, project string, serverUrl string) *Rancher {
	v := new(Rancher)
	v.homeDirectory = homeDirectory
	v.setRancherConfig(token, project, serverUrl)
	return v
}

func (r *Rancher) setRancherConfig(token string, project string, serverUrl string) {
	splitToken := strings.Split(token, ":")
	config := config{
		CurrentServer: "PrimaryServer",
		Servers: serversConfig{
			PrimaryServer: serverConfig{
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
	} else {
		ioutil.WriteFile(r.homeDirectory + "/.rancher/cli.json", jsonByteArray, 0644)
	}
}

func (r *Rancher) RunCommand(command string) string  {
	cmd := exec.Command("sh", "-c", "rancher " + command)
	cmd.Env = append(os.Environ(),
		"HOME=" + r.homeDirectory,
	)
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(output[:])
}

func (r *Rancher) RunKubectlCommand(command string) string  {
	return r.RunCommand("kubectl " + command)
}