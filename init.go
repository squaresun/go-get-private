package main

import (
	"flag"
	"io"
	"os"
)

var outWriter io.Writer

var repos []string

func init() {
	flag.Parse()
	if *verbose == true {
		outWriter = os.Stdout
	}

	repos = flag.Args()
}
