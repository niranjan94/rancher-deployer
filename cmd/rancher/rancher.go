package rancher

import (
	"strings"
	"fmt"
	"gopkg.in/resty.v1"
)

//ServerConfig holds the config for each server the user has setup
type ServerConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	TokenKey  string `json:"tokenKey"`
	URL       string `json:"url"`
	Project   string `json:"project"`
	CACerts   string `json:"cacert"`
}

type Rancher struct {
	config ServerConfig
	K8sClient *K8s
	client *resty.Client
}

type KubeConfigApiResponse struct {
	Config string `json:"config"`
}

func New(token string, project string, serverUrl string) *Rancher {
	r := new(Rancher)
	splitToken := strings.Split(token, ":")
	r.config = ServerConfig{
		AccessKey: splitToken[0],
		SecretKey: splitToken[1],
		TokenKey:  token,
		URL:       serverUrl,
		Project:   project,
		CACerts:    "",
	}
	r.client = resty.New()
	r.client.SetBasicAuth(r.config.AccessKey, r.config.SecretKey)
	r.client.HostURL = r.config.URL
	r.client.SetHeader("Accept", "application/json")

	resp, err := r.client.R().SetResult(KubeConfigApiResponse{}).Post(fmt.Sprintf("/v3/clusters/%s?action=generateKubeconfig", strings.Split(project, ":")[0]))

	if err != nil {
		panic(err.Error())
	}

	kubeConfig := resp.Result().(*KubeConfigApiResponse).Config
	r.K8sClient = NewK8sClient(kubeConfig, false)
	return r
}