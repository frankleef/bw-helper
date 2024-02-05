package cmd

import (
	"context"

	"github.com/frankleef/bw-helper/internal/session"
	"github.com/frankleef/bw-helper/internal/validator"
	"github.com/urfave/cli/v3"
)

var (
	loginCmd = &cli.Command{
		Name:        "login",
		Description: "Login to your Bitwarden Vault",
		Action: func(ctx context.Context, c *cli.Command) error {

			return executeLogin(c)
		},
	}
)

func executeLogin(c *cli.Command) error {
	if err := validator.ValidateInit(); err != nil {
		errorLogger.Fatal(err)
	}

	if err := session.InitializeSession(); err != nil {
		errorLogger.Fatal(err)
	}

	return nil
}
