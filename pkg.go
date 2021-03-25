package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/mholt/archiver/v3"
	"gopkg.in/yaml.v2"
)

type pkg struct {
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	Description  string   `yaml:"description"`
	Dependencies []string `yaml:"dependencies"`
	Files        []struct {
		Url        string
		Buildstyle string
	} `yaml:"files"`
	Patches []string `yaml:"patches"`

	isInstalled bool
	isUpToDate  bool
	path        string
}

func (p *pkg) fill(name string) error {
	path, err := searchforpackage(name)
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/build.yml", path))
	if err != nil {
		return err
	}

	tmpl, err := template.New(name).Parse(string(file))
	if err != nil {
		return err
	}

	var buff bytes.Buffer

	err = tmpl.Execute(&buff, nil)
	if err != nil {
		return err
	}

	p.path = path

	if err = ioutil.WriteFile(fmt.Sprintf("%s/build.compiled.yml", path), buff.Bytes(), 0644); err != nil {
		return err
	}

	return yaml.Unmarshal(buff.Bytes(), p)
}

func (p *pkg) installpackage() error {
	fmt.Printf("Building %s-%s...\n", p.Name, p.Version)

	if err := os.RemoveAll(fmt.Sprintf("%s/workdir", p.path)); err != nil {
		return err
	}

	if err := os.Mkdir(fmt.Sprintf("%s/workdir", p.path), 0711); err != nil {
		return err
	}

	for _, v := range p.Files {
		fmt.Printf("Downloading %s...\n", v.Url)
		if err := download(v.Url, fmt.Sprintf("%s/workdir/%s", p.path, filepath.Base(v.Url))); err != nil {
			return err
		}

		fmt.Printf("Unpacking %s...\n", filepath.Base(v.Url))
		if err := archiver.Unarchive(fmt.Sprintf("%s/workdir/%s", p.path, filepath.Base(v.Url)), fmt.Sprintf("%s/workdir", p.path)); err != nil {
			return err
		}
	}

	return nil
}
