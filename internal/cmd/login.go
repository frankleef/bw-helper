package cmd

import (
	"context"

	"github.com/frankleef/bw-helper/internal/config"
	"github.com/frankleef/bw-helper/internal/session"
	"github.com/frankleef/bw-helper/internal/validator"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
)

func addLogin() *cli.Command {
	return &cli.Command{
		Name:        "login",
		Description: "Login to your Bitwarden Vault",
		Action: func(ctx context.Context, c *cli.Command) error {

			return executeLogin(c)
		},
	}
}

func executeLogin(c *cli.Command) error {
	if err := validator.ValidateInit(); err != nil {
		errorLogger.Fatal(err)
	}

	err := viper.ReadInConfig()
	if err != nil {
		errorLogger.Fatal(err)
	}

	err = viper.Unmarshal(&config.Configuration)

	if err != nil {
		errorLogger.Fatal(err)
	}

	if err := session.InitializeSession(); err != nil {
		errorLogger.Fatal(err)
	}

	return nil
}
