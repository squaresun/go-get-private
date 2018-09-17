package main

import (
	"fmt"
	"io"
	"os"
)

// Output txt on console
type writer interface {
	io.Writer
}

// checkIfError should be used to naively panics if an error is not nil.
func checkIfError(w writer, err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// info should be used to describe the example commands that are about to run.
func info(w writer, format string, args ...interface{}) {
	w.Write([]byte(fmt.Sprintf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))))
}
