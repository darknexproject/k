package main

import (
	"os"

	"github.com/mkideal/cli"
)

var install = &cli.Command{
	Name: "install",
	Desc: "install a package",
	Fn: func(ctx *cli.Context) error {
		needroot()

		for _, v := range os.Args[2:] {
			var p pkg

			if err := p.fill(v); err != nil {
				return err
			}

			if err := p.installpackage(); err != nil {
				return err
			}
		}

		return nil
	},
}
