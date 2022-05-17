package commands_develop

import (
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"strings"
)

func DevelopTestCmd(ui cli.Ui) (cli.Command, error) {
	s := &DevelopTestCommand{}
	s.init()
	return s, nil
}

type DevelopTestCommand struct {
	flags      *flag.FlagSet
	configPath string
	helpText   string
}

func (s *DevelopTestCommand) init() {
	s.flags = flag.NewFlagSet("", flag.ContinueOnError)
	s.flags.StringVar(&s.configPath, "config", "config.yml", "Yml config file path")
	//s.helpText = flags.Usage(s.help(), s.flags)
}

func (s *DevelopTestCommand) help() string {
	return `
Usage: [options]

	Develop test command
`
}

func (s *DevelopTestCommand) Help() string {
	return s.helpText
}

func (s *DevelopTestCommand) Synopsis() string {
	return DevelopTest
}

func (s *DevelopTestCommand) Run(args []string) int {
	if err := s.flags.Parse(args); err != nil {
		if !strings.Contains(err.Error(), "help requested") {
			log.Printf("error parsing flags: %v", err)
		}
		return 1
	}
	fmt.Println("Command launch " + s.Synopsis())
	return 0
}
