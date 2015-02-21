package main

import (
	"fmt"
	"os"

	"github.com/samalba/dockerclient"
)

type serverOptions struct {
	MCDir string `long:"mcdir" description:"Directory path to materials commons file storage"`
	Name  string `long:"name" description:"Materials Commons instance name"`
}

type startCommand struct {
}

type restartCommand struct {
}

type stopCommand struct {
}

type deployCommand struct {
}

type upgradeCommand struct {
}

type Options struct {
	Server serverOptions `group:"Server Options"`
}

var (
	docker dockerclient.Client
)

func main() {
	docker, err := dockerclient.NewDockerClient("", nil)
	if err != nil {
		fmt.Printf("Unable to connect to docker: %s", err)
		os.Exit(1)
	}
	var _ = docker
}

func (cmd *startCommand) Execute(args []string) error {
	return nil
}

func (cmd *restartCommand) Execute(args []string) error {
	return nil
}

func (cmd *stopCommand) Execute(args []string) error {
	return nil
}

func (cmd *deployCommand) Execute(args []string) error {
	return nil
}

func (cmd *upgradeCommand) Execute(args []string) error {
	return nil
}
