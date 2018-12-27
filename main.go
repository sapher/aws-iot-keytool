package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

type options struct {
}

var opts options

var parser = flags.NewParser(&opts, flags.Default)

func main() {
	exitCommandNotFound("openssl")
	exitCommandNotFound("keytool")

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
