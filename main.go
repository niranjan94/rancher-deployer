package main

import (
	"log"
	"os"
	"github.com/urfave/cli"
	"github.com/niranjan94/rancher-deployer/actions"
)

func main() {
	app := cli.NewApp()
	app.Name = "rancher-deployer"
	app.Usage = "Deploy/upgrade your deployments on Rancher 2.0 Clusters"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "token",
			Value: "",
			Usage: "Override token for deployment",
		},
		cli.StringFlag{
			Name: "tag",
			Value: "",
			Usage: "Override the tag for the docker image to use",
		},
		cli.StringFlag{
			Name: "environments",
			Value: "",
			Usage: "Environments to deploy to (comma-separated)",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("token") != "" {
			os.Setenv("DEPLOYER_TOKEN", c.String("token"))
		}
		actions.LoadConfig()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
