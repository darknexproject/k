package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Repositories []struct {
		Url  string `yaml:"url"`
		Path string `yaml:"localpath"`
	} `yaml:"repositories"`
}

func (c *config) fill() error {
	file, err := ioutil.ReadFile("/k/config.yml")
	if err != nil {
		return err
	}

	return yaml.Unmarshal(file, c)
}

func (c *config) update() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("/k/config.yml", data, 0644)
}
