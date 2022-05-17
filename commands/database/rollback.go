package commands_database

import (
	"flag"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/kachit/golang-api-skeleton/bootstrap"
	"github.com/kachit/golang-api-skeleton/database/migrations"
	"github.com/mitchellh/cli"
	"log"
	"strings"
)

func RollbackCmd(ui cli.Ui) (cli.Command, error) {
	s := &Rollback{}
	s.init()
	return s, nil
}

type Rollback struct {
	flags      *flag.FlagSet
	configPath string
	migrateId  string
	helpText   string
}

func (s *Rollback) init() {
	s.flags = flag.NewFlagSet("", flag.ContinueOnError)
	s.flags.StringVar(&s.configPath, "config", "config.yml", "Yml config file path")
	s.flags.StringVar(&s.migrateId, "migrateId", "", "migration id")
}

func (s *Rollback) help() string {
	return `
Usage: [options]

	Rollback migrations.
`
}

func (s *Rollback) Help() string {
	return s.helpText
}

func (s *Rollback) Synopsis() string {
	return `rollback`
}

func (s *Rollback) Run(args []string) int {
	if err := s.flags.Parse(args); err != nil {
		if !strings.Contains(err.Error(), "help requested") {
			log.Printf("error parsing flags: %v", err)
		}
		return 1
	}

	err := s.runMigrations()
	if err != nil {
		log.Println(err.Error())
		return 1
	}
	return 0
}

func (s *Rollback) runMigrations() error {
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
		if err = m.RollbackTo(s.migrateId); err != nil {
			return err
		}
	} else {
		if err = m.RollbackLast(); err != nil {
			return err
		}
	}
	return nil
}
