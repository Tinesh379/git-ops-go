package main

import (
	"os"
)

const (
	repoURL        string = "https://github.com/Tinesh379/webapp.git"
	cloneDirectory string = "D:\\UpSkilling\\golang\\workdir\\webapp"
)

var (
	gitUser string
	gitPass string
)

func init() {
	// Load GITHUB user and password from env variables
	gitUser = os.Getenv("GITHUB_USER")
	gitPass = os.Getenv("GITHUB_PAT")
}

func main() {
	// name of the file to make changes
	var fileName string = "README.md"
	GitClone()
	EditFile(fileName)
	GitAddAndCommit(fileName)
	GitPush()

}
