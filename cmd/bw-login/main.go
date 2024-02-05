package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"

	"github.com/frankleef/bw-login/internal/config"
	"github.com/frankleef/bw-login/internal/session"
	"github.com/frankleef/bw-login/internal/validator"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
)

func main() {
	infoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	errorLogger := log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile)

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("$HOME/.bw-helper")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		infoLogger.Println("config file was not found in $HOME/.bw-helper")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		errorLogger.Fatal(err)
	}

	config.Configuration = config.Config{HomeDir: home, Scheme: "http", Host: "localhost", Port: 8087, ConfigDir: ".bw-helper"}

	cmd := &cli.Command{
		Commands: []*cli.Command{
			&cli.Command{
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
					infoLogger.Println("trying to create config dir...")
					if err := os.Mkdir(fmt.Sprintf("%s/%s", config.Configuration.HomeDir, config.Configuration.ConfigDir), os.ModePerm); err != nil {
						return err
					}

					if err := config.Configuration.InitConfig(c.String("scheme"), c.String("host"), c.Int("port"), c.String("password")); err != nil {
						errorLogger.Fatal(err)
					}

					infoLogger.Println("config dir created. Run `bw-login login`.")
					return nil
				},
			},
			&cli.Command{
				Name:        "login",
				Description: "Login to your Bitwarden Vault",
				Action: func(ctx context.Context, c *cli.Command) error {

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
				},
			},
			&cli.Command{
				Name:        "reset",
				Description: "Reset your password",
				Action: func(ctx context.Context, c *cli.Command) error {
					fmt.Println("Reset")
					return nil
				},
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {

			if err := validator.ValidateInit(); err != nil {
				errorLogger.Fatal(err)
			}

			session.SetSession()
			args := c.Args().Slice()

			bwCmd := exec.Command("bw", args[:]...)

			var outb bytes.Buffer
			bwCmd.Stdout = &outb
			bwCmd.Stderr = os.Stderr

			err := bwCmd.Run()

			if err != nil {
				errorLogger.Fatal(err)
			}

			if c.Args().First() == "get" {
				infoLogger.Println("Copying output to clipboard")
				if err := clipboard.WriteAll(outb.String()); err != nil {
					errorLogger.Fatal(err)
				}
				infoLogger.Println("Success")
				return nil
			}

			infoLogger.Println(outb.String())
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
