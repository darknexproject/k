package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}

func needroot() {
	if os.Geteuid() != 0 {
		check(errors.New("Man, this task requires root access!"))
	}
}

func newprefix(path string) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 32)

	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	res := string(b)

	if _, err := os.Stat(fmt.Sprintf("%s/%s", path, res)); os.IsExist(err) {
		return newprefix(path)
	}

	return fmt.Sprintf("%s/%s", path, res)
}
