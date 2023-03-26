package develop

import (
	"fmt"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/urfave/cli/v2"
)

func NewDevelopTestCommand() *cli.Command {
	return &cli.Command{
		Name:  "develop:test",
		Usage: "Test drive command",
		Flags: []cli.Flag{},
		Action: func(cCtx *cli.Context) error {
			config, err := infrastructure.NewConfig("config.yml")
			if err != nil {
				return err
			}
			fmt.Println(config)
			return nil
		},
	}
}
