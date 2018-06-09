package main

import (
	"log"
	"os"
	"github.com/urfave/cli"
	"github.com/niranjan94/rancher-deployer/cmd"
	"fmt"
)

var version = "snapshot"
var revision = "head"

func main() {

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s\nversion=%s revision=%s\n",c.App.Name, c.App.Version, revision)
	}

	app := cli.NewApp()
	app.Name = "rancher-deployer"
	app.Usage = "Deploy/upgrade your deployments on Rancher 2.0 Clusters"
	app.Version = version

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
			Name: "image",
			Value: "",
			Usage: "Override the docker image",
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
		cmd.LoadConfig()
		cmd.Deploy(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

