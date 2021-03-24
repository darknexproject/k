package main

import "github.com/mkideal/cli"

var install = &cli.Command{
	Name: "install",
	Desc: "install a package",
	Fn: func(ctx *cli.Context) error {
		needroot()


		return nil
	}
}