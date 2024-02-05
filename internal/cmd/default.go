package cmd

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/frankleef/bw-helper/internal/config"
	"github.com/frankleef/bw-helper/internal/session"
	"github.com/urfave/cli/v3"
)

func defaultCmd(c *cli.Command) error {
	if err := config.Configuration.ValidateConfig(); err != nil {
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
}
