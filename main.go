package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mkideal/cli"
)

var cfg config
var db database
var help = cli.HelpCommand("print this message")

var root = &cli.Command{
	Desc: "Based package manager for Karoshi Linux",
	Fn: func(c *cli.Context) error {
		return errors.New(fmt.Sprintf("Hey buddy, I think you forgot something, try \"%s help\"!", os.Args[0]))
	},
}

func main() {
	rand.Seed(time.Now().UnixNano())
	check(cfg.fill())
	check(db.fill())

	check(cli.Root(
		root,
		cli.Tree(help),
		cli.Tree(sync),
	).Run(os.Args[1:]))
}
