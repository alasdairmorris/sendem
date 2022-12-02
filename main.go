package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `A command-line tool for sending emails via SMTP.

Usage:
  sendem
  sendem -h | --help
  sendem --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.
`

	opts, _ := docopt.ParseArgs(usage, nil, "https://github.com/alasdairmorris/sendem v0.0.1")
	fmt.Println(opts)
}
