package validator

import (
	"errors"
	"fmt"
	"os"

	"github.com/frankleef/bw-helpers/internal/config"
)

func ValidateInit() error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", config.Configuration.HomeDir, config.Configuration.ConfigDir)); err != nil {
		return errors.New("Directory $HOME/.bw-helper does not exist. Run bw-login init first.")
	}

	return nil
}