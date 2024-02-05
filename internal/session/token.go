package session

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/frankleef/bw-helpers/internal/vault"
)

func GetSessionToken() error {
	_, err := exec.LookPath("bw")

	if err != nil {
		return errors.New("Bitwarden CLI is not installed or available in $PATH.")
	}

	bwCmd := exec.Command("bw", "serve")
	bwCmd.Stdout = os.Stdout
	if err := bwCmd.Start(); err != nil {
		return errors.New("Could not start Bitwarden Vault Management API")
	}
	time.Sleep(3 * time.Second)
	if code, err := vault.Unlock(); err != nil {
		return err
	} else {
		fmt.Println(*code)
	}

	if err != nil {
		return err
	}

	bwCmd.Process.Kill()
	return nil
}
