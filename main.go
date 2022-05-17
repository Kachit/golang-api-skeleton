package main

import (
	"fmt"
	"github.com/kachit/golang-api-skeleton/commands"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

func main() {
	c := cli.NewCLI("golang-api-skeleton", "1.0.0")
	c.Args = os.Args[1:]
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	c.Commands = commands.GetCmds(ui)

	fmt.Println(c)
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
