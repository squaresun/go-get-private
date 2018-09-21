package main

import (
	"flag"

	git "gopkg.in/src-d/go-git.v4"
)

var remoteName = "origin"

var update = flag.Bool("u", false, "Update existing git repo")

func updateRepo(proj *git.Repository) error {
	tree, err := proj.Worktree()
	if err != nil {
		return err
	}
	// Pull the latest changes from the origin remote and merge into the current branch
	remote, err := proj.Remote(remoteName)
	checkIfError(err)
	info(outWriter, "git pull %s", remote.Config().URLs[0])
	err = tree.Pull(&git.PullOptions{
		RemoteName: remoteName,
		Progress:   outWriter,
	})
	if err != nil {
		if err != git.NoErrAlreadyUpToDate {
			return err
		}
		info(outWriter, err.Error())
		return nil
	}

	// Print the latest commit that was just pulled
	ref, err := proj.Head()
	if err != nil {
		return err
	}

	commit, err := proj.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	info(outWriter, commit.String())

	return nil
}
