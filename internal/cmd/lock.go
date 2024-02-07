package cmd

import (
	"context"
	"os"
	"os/exec"

	"github.com/frankleef/bw-helper/internal/session"
	"github.com/urfave/cli/v3"
)

var (
	lockCmd = &cli.Command{
		Name:        "lock",
		Description: "Lock your Bitwarden Vault",
		Action: func(ctx context.Context, c *cli.Command) error {

			return executeLock(c)
		},
	}
)

func executeLock(c *cli.Command) error {
	if err := session.RemoveSession(); err != nil {
		errorLogger.Fatal(err)
	}

	bwCmd := exec.Command("bw", "lock")

	bwCmd.Stdout = os.Stdout
	bwCmd.Stderr = os.Stderr
	bwCmd.Stdin = os.Stdin

	if err := bwCmd.Run(); err != nil {
		return err
	}

	return nil
}
