package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Configuration = Config{}

type Config struct {
	HomeDir   string `yaml:"-"`
	ConfigDir string `yaml:"-"`
	Host      string `yaml:"host"`
	Port      int64  `yaml:"port"`
	Password  string `yaml:"password"`
}

func (c *Config) InitConfig(host string, port int64, password string) error {
	if host != "" {
		c.Host = host
	}

	if port != 0 {
		c.Port = port
	}

	c.Password = password

	data, err := yaml.Marshal(&c)

	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/.bw-helper/config.yaml", c.HomeDir), data, 0644)

	if err != nil {
		return err
	}

	return nil
}
