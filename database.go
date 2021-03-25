package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type databaseRepo struct {
	Url  string `yaml:"url"`
	Path string `yaml:"localpath"`
}

type databasePackage struct {
	Name    string   `yaml:"name"`
	Prefix  string   `yaml:"prefix"`
	Version string   `yaml:"version"`
	Files   []string `yaml:"files"`
}

type database struct {
	Packages     []databasePackage `yaml:"packages"`
	Repositories []databaseRepo    `yaml:"repositories"`
}

func (d *database) fill() error {
	file, err := ioutil.ReadFile("/k/db.yml")
	if err != nil {
		return nil
	}

	return yaml.Unmarshal(file, d)
}

func (d *database) update() error {
	data, err := yaml.Marshal(d)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("/k/db.yml", data, 0644)
}

func (d *database) getrepopath(url string) (string, error) {
	for _, v := range d.Repositories {
		if v.Url == url {
			return v.Path, nil
		}
	}

	return "", errors.New("No such repository")
}

func (d *database) addRepo(url string, path string) {
	d.Repositories = append(d.Repositories, databaseRepo{url, path})
}
