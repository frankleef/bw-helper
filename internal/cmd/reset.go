package cmd

import (
	"context"

	"github.com/frankleef/bw-helper/internal/config"
	"github.com/frankleef/bw-helper/internal/validator"
	"github.com/urfave/cli/v3"
)

var (
	resetCmd = &cli.Command{
		Name:        "reset",
		Description: "Reset your password",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "password",
				Usage:    "Password to login into Bitwarden",
				Required: true,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return executeReset(c)
		},
	}
)

func executeReset(c *cli.Command) error {
	if err := validator.ValidateInit(); err != nil {
		errorLogger.Fatal(err)
	}

	if err := config.Configuration.UpdatePassword(c.String("password")); err != nil {
		errorLogger.Fatal(err)
	}

	infoLogger.Println("password updated")
	return nil
}
