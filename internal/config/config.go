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
	Scheme    string `yaml:"scheme"`
	Host      string `yaml:"host"`
	Port      int64  `yaml:"port"`
	Password  string `yaml:"password"`
}

func (c *Config) InitConfig(scheme string, host string, port int64, password string) error {
	if scheme != "" {
		c.Scheme = scheme
	}

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

func (c *Config) UpdatePassword(password string) error {
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
