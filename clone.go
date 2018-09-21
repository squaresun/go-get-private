package main

import (
	"os"
	"path"
	"regexp"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

// Handles Both grandpeak.com and ssh://git@192.168.135.218:8822
var urlMap = map[string]string{
	"grandpeak.com": "ssh://git@192.168.135.218:8822",
	// Handling if you directly copied from gitea url
	"ssh://git@192.168.135.218:8822": "ssh://git@192.168.135.218:8822",
}

var repoNameMap = map[string]string{
	"grandpeak.com":                  "grandpeak.com",
	"ssh://git@192.168.135.218:8822": "grandpeak.com",
}

// Support - and _
var repoReg = regexp.MustCompile("/[a-zA-Z0-9\\-\\_]+/")

func clone(repo string) (*git.Repository, error) {
	repoName := repoReg.Split(repo, -1)[0]
	repoPath := strings.Replace(repo, repoName, repoNameMap[repoName], -1)
	gitPath := path.Join(os.Getenv("GOPATH"), "src", repoPath)
	proj, err := git.PlainOpen(gitPath)
	if err == git.ErrDestinationExists {
		return proj, nil
	}
	if err == git.ErrRepositoryNotExists {
		urlPath := strings.Replace(repo, repoName, urlMap[repoName], -1)
		info(outWriter, "git clone %s", gitPath)
		proj, err = git.PlainClone(gitPath, false, &git.CloneOptions{
			URL:      urlPath,
			Progress: outWriter,
		})
	}
	return proj, err
}
