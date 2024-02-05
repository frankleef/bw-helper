package validator

import (
	"errors"

	"github.com/frankleef/bw-helper/internal/config"
)

func ValidateInit() error {
	if !config.Configuration.ConfigExists() {
		return errors.New("directory $HOME/.bw-helper does not exist. Run bw-helper init first")
	}

	return nil
}
