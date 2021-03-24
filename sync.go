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

		for _, x := range cfg.Repositories {
			fmt.Printf("Fetching %s...\n", x.Url)

			var err error

			path, err := db.getrepopath(x.Url)
			if err == nil {
				r, err := git.PlainOpen(path)
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

				db.addRepo(x.Url, prefix)
				db.update()
			}

			if err != nil {
				return err
			}
		}

		return nil
	},
}
