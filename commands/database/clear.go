package commands_database

import (
	"flag"
	"fmt"
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/kachit/golang-api-skeleton/database/seeds"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"github.com/mitchellh/cli"
	"log"
	"strings"
)

func ClearCmd(ui cli.Ui) (cli.Command, error) {
	s := &ClearSeed{}
	s.init()
	return s, nil
}

type ClearSeed struct {
	flags      *flag.FlagSet
	configPath string
	confirm    bool
	helpText   string
}

func (s *ClearSeed) init() {
	s.flags = flag.NewFlagSet("", flag.ContinueOnError)
	s.flags.StringVar(&s.configPath, "config", "config.yml", "Yml config file path")
	s.flags.BoolVar(&s.confirm, "confirm", false, "Confirmation flag")
}

func (s *ClearSeed) help() string {
	return `
Usage: [options]

	Starts database clear.
`
}

func (s *ClearSeed) Help() string {
	return s.helpText
}

func (s *ClearSeed) Synopsis() string {
	return `clear`
}

func (s *ClearSeed) Run(args []string) int {
	if err := s.flags.Parse(args); err != nil {
		if !strings.Contains(err.Error(), "help requested") {
			log.Printf("error parsing flags: %v", err)
		}
		return 1
	}
	fmt.Println("Command launch " + s.Synopsis())

	err := s.runClear()
	if err != nil {
		log.Println(err.Error())
		return 1
	}
	return 0
}

func (s *ClearSeed) runClear() error {
	container, err := bootstrap.InitializeContainer(s.configPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	seedersStack := gorm_seeder.NewSeedersStack(container.DB)
	if s.confirm {
		config := gorm_seeder.SeederConfiguration{}
		usersSeeder := seeds.NewUsersSeeder(config)
		seedersStack.AddSeeder(&usersSeeder)
	}
	err = seedersStack.Clear()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
