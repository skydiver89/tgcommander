package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Telegram Telegram `yaml:"telegram"`
	Buttons  []Button `yaml:"buttons"`
}

type Telegram struct {
	Token          string  `yaml:"token"`
	Users          []int64 `yaml:"users"`
	DeclineMessage string  `yaml:"declineMessage"`
}

type Button struct {
	Name     string `yaml:"name"`
	Row      int    `yaml:"row"`
	Command  string `yaml:"command"`
	Script   string `yaml:"script"`
	Output   bool   `yaml:"output"`
	ExitCode bool   `yaml:"exitCode"`
}

func (c *Config) load(fname string) error {
	yamlData, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlData, c)
	return err
}
