package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
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

func (c *Config) InitConfig() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("$HOME/.bw-helper")
	viper.ReadInConfig() // Find and read the config file

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	Configuration = Config{HomeDir: home, Scheme: "http", Host: "localhost", Port: 8087, ConfigDir: ".bw-helper"}

	if !c.ConfigExists() {
		return nil
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}

func (c *Config) SetConfig(scheme string, host string, port int64, password string) error {
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

func (c *Config) ConfigExists() bool {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", c.HomeDir, c.ConfigDir)); err != nil {
		return false
	}

	return true
}
