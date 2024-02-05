package cmd

import (
	"context"
	"log"
	"os"

	"github.com/frankleef/bw-helper/internal/config"
	"github.com/urfave/cli/v3"
)

var infoLogger = *log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
var errorLogger = *log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile)

func Execute() {

	if err := config.Configuration.InitConfig(); err != nil {
		errorLogger.Fatal(err)
	}

	cmd := &cli.Command{
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			initCmd,
			loginCmd,
			resetCmd,
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return defaultCmd(c)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
