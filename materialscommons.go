package main

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
	"github.com/jessevdk/go-flags"
)

type StartCommand struct {
	Name string `long:"name" description:"Materials Commons instance name"`
}

type RestartCommand struct {
	Name string `long:"name" description:"Materials Commons instance name"`
}

type StopCommand struct {
	Name string `long:"name" description:"Materials Commons instance name"`
}

type DeployCommand struct {
	MCDir string `long:"mcdir" description:"Directory path to materials commons file storage"`
	Name  string `long:"name" description:"Materials Commons instance name"`
}

type UpgradeCommand struct {
}

var (
	client         docker.Client
	startCommand   StartCommand
	restartCommand RestartCommand
	stopCommand    StopCommand
	deployCommand  DeployCommand
	upgradeCommand UpgradeCommand
)

func main() {
	parser := flags.NewParser(nil, flags.Default)
	_, err := parser.Parse()
	parser.AddCommand("start", "Start a MaterialsCommons instance", "", &startCommand)
	parser.AddCommand("restart", "Restarts a MaterialsCommons instance", "", &restartCommand)
	parser.AddCommand("stop", "Stops a MaterialsCommons instance", "", &stopCommand)
	parser.AddCommand("deploy", "Deploys a new MaterialsCommons instance", "", &deployCommand)
	parser.AddCommand("upgrade", "Upgrades a MaterialsCommons instance", "", &upgradeCommand)

	if err != nil {
		fmt.Printf("Failed parsing commands: %s\n", err)
		os.Exit(1)
	}

	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		fmt.Printf("Failed to connect to docker: %s\n", err)
		os.Exit(1)
	}

	var _ = client

	fmt.Println("Connected to docker!")

	// docker, err := dockerclient.NewDockerClient("", nil)
	// if err != nil {
	// 	fmt.Printf("Unable to connect to docker: %s", err)
	// 	os.Exit(1)
	// }
	// var _ = docker
}

func (cmd *StartCommand) Execute(args []string) error {
	return nil
}

func (cmd *RestartCommand) Execute(args []string) error {
	return nil
}

func (cmd *StopCommand) Execute(args []string) error {
	return nil
}

func (cmd *DeployCommand) Execute(args []string) error {
	return nil
}

func (cmd *UpgradeCommand) Execute(args []string) error {
	return nil
}
