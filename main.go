package main

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

func main() {

	info("git clone https://github.com/src-d/go-git")
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})

	checkIfError(err)
}
