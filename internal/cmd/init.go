package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/frankleef/bw-helper/internal/config"
	"github.com/urfave/cli/v3"
)

func addInit() *cli.Command {
	return &cli.Command{
		Name:        "init",
		Description: "Initialize the helper.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "password",
				Usage:    "Password to login into Bitwarden",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "scheme",
				Usage: "Scheme of Vault Management API. Default http",
			},
			&cli.StringFlag{
				Name:  "host",
				Usage: "Host of Vault Management API. Default localhost",
			},
			&cli.StringFlag{
				Name:  "port",
				Usage: "Vault Management API port. Default 8087",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return executeInit(c)
		},
	}
}

func executeInit(c *cli.Command) error {
	infoLogger.Println("trying to create config dir...")
	if err := os.Mkdir(fmt.Sprintf("%s/%s", config.Configuration.HomeDir, config.Configuration.ConfigDir), os.ModePerm); err != nil {
		return err
	}

	if err := config.Configuration.SetConfig(c.String("scheme"), c.String("host"), c.Int("port"), c.String("password")); err != nil {
		errorLogger.Fatal(err)
	}

	infoLogger.Println("config dir created. Run `bw-helper login`.")
	return nil
}
