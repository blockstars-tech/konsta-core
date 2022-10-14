package main

import (
	_ "embed"

	"konsta.live/command/root"
	"konsta.live/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
