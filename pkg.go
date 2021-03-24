package main

type pkg struct {
	Name         string
	Version      string
	Description  string
	Dependencies []string
	Files        []struct {
		Url        string
		Buildstyle string
	}
	Patches []string

	isInstalled bool
}

func (p *pkg) fill(name string) {

}
