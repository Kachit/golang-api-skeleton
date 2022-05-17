package commands_database

import (
	"flag"
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/kachit/golang-api-skeleton/database/migrations"
	"github.com/mitchellh/cli"
	"log"
	"strings"
)

func MigrateCmd(ui cli.Ui) (cli.Command, error) {
	s := &Migrate{}
	s.init()
	return s, nil
}

type Migrate struct {
	flags      *flag.FlagSet
	configPath string
	migrateId  string
	helpText   string
}

func (s *Migrate) init() {
	s.flags = flag.NewFlagSet("", flag.ContinueOnError)
	s.flags.StringVar(&s.configPath, "config", "config.yml", "Yml config file path")
	s.flags.StringVar(&s.migrateId, "migrateId", "", "migrate id")
}

func (s *Migrate) help() string {
	return `
Usage: [options]

	Starts migrations.
`
}

func (s *Migrate) Help() string {
	return s.helpText
}

func (s *Migrate) Synopsis() string {
	return `migrate`
}

func (s *Migrate) Run(args []string) int {
	if err := s.flags.Parse(args); err != nil {
		if !strings.Contains(err.Error(), "help requested") {
			log.Printf("error parsing flags: %v", err)
		}
		return 1
	}
	fmt.Println("Command launch " + s.Synopsis())

	err := s.runMigrations()
	if err != nil {
		log.Println(err.Error())
		return 1
	}
	return 0
}

func (s *Migrate) runMigrations() error {
	container, err := bootstrap.InitializeContainer(s.configPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	m := gormigrate.New(
		container.DB,
		gormigrate.DefaultOptions,
		migrations.Migrations,
	)
	if s.migrateId != "" {
		if err = m.MigrateTo(s.migrateId); err != nil {
			return err
		}
	} else {
		if err = m.Migrate(); err != nil {
			return err
		}
	}
	return nil
}
