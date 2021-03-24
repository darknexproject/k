package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/mkideal/cli"
)

var sync = &cli.Command{
	Name: "sync",
	Desc: "synchronise repos",
	Fn: func(ctx *cli.Context) error {
		needroot()

		for i, x := range cfg.Repositories {
			fmt.Printf("Fetching %s...\n", x.Url)

			var err error

			if x.Path != "" {
				r, err := git.PlainOpen(x.Path)
				if err != nil {
					return err
				}

				w, err := r.Worktree()
				if err != nil {
					return err
				}

				err = w.Pull(&git.PullOptions{
					RemoteName: "origin",
				})
			} else {
				prefix := newprefix("/k/repositories")
				_, err = git.PlainClone(prefix, false, &git.CloneOptions{
					URL: x.Url,
				})

				cfg.Repositories[i].Path = prefix
				cfg.update()
			}

			if err != nil {
				return err
			}
		}

		return nil
	},
}
