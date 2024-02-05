package cmd

import (
	"context"
	"log"
	"os"

	"github.com/frankleef/bw-helper/internal/config"
	"github.com/urfave/cli/v3"
)

var infoLogger log.Logger
var errorLogger log.Logger

func Execute() {
	initLoggers()

	if err := config.Configuration.InitConfig(); err != nil {
		errorLogger.Fatal(err)
	}

	cmd := &cli.Command{
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			addInit(),
			addLogin(),
			addReset(),
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return passCommandToBw(c)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func initLoggers() {
	infoLogger = *log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	errorLogger = *log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile)
}
