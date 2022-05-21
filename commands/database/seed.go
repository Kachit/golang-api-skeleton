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

func SeedCmd(ui cli.Ui) (cli.Command, error) {
	s := &Seed{}
	s.init()
	return s, nil
}

type Seed struct {
	flags      *flag.FlagSet
	configPath string
	users      int
	helpText   string
}

func (s *Seed) init() {
	s.flags = flag.NewFlagSet("", flag.ContinueOnError)
	s.flags.StringVar(&s.configPath, "config", "config.yml", "Yml config file path")
	s.flags.IntVar(&s.users, "users", 10, "Users rows count")
}

func (s *Seed) help() string {
	return `
Usage: [options]

	Starts seeding.
`
}

func (s *Seed) Help() string {
	return s.helpText
}

func (s *Seed) Synopsis() string {
	return `seed`
}

func (s *Seed) Run(args []string) int {
	if err := s.flags.Parse(args); err != nil {
		if !strings.Contains(err.Error(), "help requested") {
			log.Printf("error parsing flags: %v", err)
		}
		return 1
	}
	fmt.Println("Command launch " + s.Synopsis())

	err := s.runSeeding()
	if err != nil {
		log.Println(err.Error())
		return 1
	}
	return 0
}

func (s *Seed) runSeeding() error {
	container, err := bootstrap.InitializeContainer(s.configPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//db := container.DB.Ses
	seedersStack := gorm_seeder.NewSeedersStack(container.DB)
	if s.users > 0 {
		cfg := gorm_seeder.SeederConfiguration{Rows: s.users}
		usersSeeder := seeds.NewUsersSeeder(cfg, container.PG)
		seedersStack.AddSeeder(&usersSeeder)
	}
	err = seedersStack.Seed()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
