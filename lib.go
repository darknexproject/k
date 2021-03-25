package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
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

func searchforpackage(name string) (string, error) {
	m, err := filepath.Glob(fmt.Sprintf("/k/repositories/*/%s", name))
	if len(m) == 0 {
		return "", errors.New(fmt.Sprintf("Package %s not found", name))
	}

	return m[0], err
}

func download(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func trimext(filename string) string {
	var ret []rune
	runes := []rune(filename)

	for _, v := range runes {
		if v == '.' {
			break
		}

		ret = append(ret, v)
	}

	return string(ret)
}
