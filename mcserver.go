package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
)

var (
	StartCommand = cli.Command{
		Name:   "start",
		Usage:  "Start the server",
		Action: startCLI,
	}

	StopCommand = cli.Command{
		Name:   "stop",
		Usage:  "Stop the server",
		Action: stopCLI,
	}

	StatusCommand = cli.Command{
		Name:   "status",
		Usage:  "Status of server",
		Action: statusCLI,
	}

	RestartCommand = cli.Command{
		Name:   "restart",
		Usage:  "Restart the server",
		Action: restartCLI,
	}

	DeployCommand = cli.Command{
		Name:   "deploy",
		Usage:  "Deploy the server",
		Action: deployCLI,
	}

	UpgradeCommand = cli.Command{
		Name:   "upgrade",
		Usage:  "Upgrade the server",
		Action: upgradeCLI,
	}
)

var (
	client docker.Client
)

func main() {
	var (
		err error
	)

	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		{
			Name:  "V. Glenn Tarcea",
			Email: "gtarcea@umich.edu",
		},
	}

	app.Commands = []cli.Command{
		StartCommand,
		StopCommand,
		StatusCommand,
		RestartCommand,
		DeployCommand,
		UpgradeCommand,
	}

	endpoint := "unix:///var/run/docker.sock"
	client, err = docker.NewClient(endpoint)
	if err != nil {
		fmt.Printf("Failed to connect to docker: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to docker!")

	app.Run(os.Args)

	// docker, err := dockerclient.NewDockerClient("", nil)
	// if err != nil {
	// 	fmt.Printf("Unable to connect to docker: %s", err)
	// 	os.Exit(1)
	// }
	// var _ = docker
}

func startCLI(context *cli.Context) {
}

func stopCLI(context *cli.Context) {
}

func statusCLI(context *cli.Context) {
}

func restartCLI(context *cli.Context) {
}

func deployCLI(context *cli.Context) {
}

func upgradeCLI(context *cli.Context) {
}
