package cmd

import (
	"github.com/urfave/cli"
	"strings"
	"github.com/spf13/viper"
	"github.com/niranjan94/rancher-deployer/cmd/utils"
	"github.com/niranjan94/rancher-deployer/cmd/rancher"
	"github.com/fatih/color"
)

func Deploy(c *cli.Context) {

	if c.String("environments") == "" {
		color.Yellow("No environments specified. Skipping deployments.")
		return
	}

	environments := strings.Split(c.String("environments"), ",")

	tagOverride := c.String("tag")
	imageOverride := c.String("image")

	globalToken := viper.GetString("token")
	globalRancherUrl := viper.GetString("rancherUrl")

	for _, environment := range environments {
		configPath := "environments." + environment
		token := viper.GetString(configPath + ".token")

		if  token == "" {
			token = globalToken
		}

		if !viper.IsSet(configPath) {
			color.Red("! Skipping invalid environment %s", environment)
			continue
		}

		for _, config := range utils.GetSliceInterfaceAsSubs(viper.Get("environments." + environment)) {
			rancherUrl := config.GetString("rancherUrl")
			project := config.GetString("project")
			image := config.GetString("image")
			deployment := config.GetString("deployment")
			namespace := config.GetString("namespace")

			if  rancherUrl == "" {
				rancherUrl = globalRancherUrl
			}

			ranch := rancher.New(token, project, rancherUrl)

			if tagOverride != "" && (image != "" || imageOverride != "") {
				imageToUse := imageOverride
				if imageOverride == "" {
					imageToUse = image
				}
				ranch.K8sClient.UpdateDeploymentImage(namespace, deployment, imageToUse, tagOverride)
			}
			ranch.K8sClient.UpdateDeployment(namespace, deployment, rancher.GetDeploymentPatchString())
			color.Green("✔Updated deployment %s in %s", deployment, namespace)
		}
	}
}
