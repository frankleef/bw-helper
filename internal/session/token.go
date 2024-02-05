package session

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/frankleef/bw-helpers/internal/config"
	"github.com/frankleef/bw-helpers/internal/vault"
)

func InitializeSession() error {
	_, err := exec.LookPath("bw")

	if err != nil {
		return errors.New("bitwarden CLI is not installed or available in $PATH")
	}

	bwCmd := exec.Command("bw", "serve")
	bwCmd.Stdout = os.Stdout
	if err := bwCmd.Start(); err != nil {
		return errors.New("could not start Bitwarden Vault Management API")
	}
	time.Sleep(3 * time.Second)
	if code, err := vault.Unlock(); err != nil {
		return err
	} else {
		if err := writeToken(*code); err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	bwCmd.Process.Kill()
	return nil
}

func SetSession() error {
	code, err := os.ReadFile(fmt.Sprintf("%s/%s/.token", config.Configuration.HomeDir, config.Configuration.ConfigDir))

	if err != nil {
		return err
	}

	os.Setenv("BW_SESSION", string(code))

	return nil
}

func writeToken(code string) error {
	file, err := os.OpenFile(fmt.Sprintf("%s/%s/.token", config.Configuration.HomeDir, config.Configuration.ConfigDir), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	if _, err := file.Write([]byte(code)); err != nil {
		file.Close()
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}
