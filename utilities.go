package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var verbose = flag.Bool("v", false, "Verbose output, i.e. print out progress")

// checkIfError should be used to naively panics if an error is not nil.
func checkIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// info should be used to describe the example commands that are about to run.
func info(w io.Writer, format string, args ...interface{}) {
	if w == nil {
		return
	}

	w.Write([]byte(fmt.Sprintf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))))
}
