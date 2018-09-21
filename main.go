package main

func main() {
	for _, repo := range repos {
		proj, err := clone(repo)
		checkIfError(err)
		if *update {
			err = updateRepo(proj)
			checkIfError(err)
		}
	}
}
