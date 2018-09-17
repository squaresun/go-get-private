package main

import (
	"os"

	"path"

	git "gopkg.in/src-d/go-git.v4"
)

func main() {
	input := "github.com/src-d/go-git"
	info(nil, "git clone https://github.com/src-d/go-git")
	_, err := git.PlainClone(path.Join(os.Getenv("GOPATH"), "src", input), false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: nil,
	})

	checkIfError(os.Stderr, err)
}
