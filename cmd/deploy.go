package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/niranjan94/rancher-deployer/cmd/rancher"
	"github.com/niranjan94/rancher-deployer/cmd/utils"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"strings"
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
			deploymentName := config.GetString("deployment")
			namespace := config.GetString("namespace")
			deploymentType := config.GetString("type")

			if deploymentType == "" {
				deploymentType = "deployment"
			}

			if  rancherUrl == "" {
				rancherUrl = globalRancherUrl
			}

			ranch := rancher.New(token, project, rancherUrl)

			id := fmt.Sprintf("%s:%s:%s", deploymentType, namespace, deploymentName)
			var err error
			if tagOverride != "" && (image != "" || imageOverride != "") {
				imageToUse := imageOverride
				if imageOverride == "" {
					imageToUse = image
				}
				err = ranch.UpdateImage(id, imageToUse, tagOverride)
			} else {
				err = ranch.Redeploy(id)
			}

			if err != nil {
				color.Green("✔Updated %s %s in %s", deploymentType, deploymentName, namespace)
			} else {
				color.Red("✗Failed to update %s %s in %s : %s", deploymentType, deploymentName, namespace, err.Error())
			}
		}
	}
}
