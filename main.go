package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func main() {

	gitUser := os.Getenv("GITHUB_USER")
	gitPass := os.Getenv("GITHUB_PAT")
	fmt.Println("Username:", gitUser)
	fmt.Println("Password:", gitPass)
	// Set the repository URL to clone
	repoURL := "https://github.com/Tinesh379/webapp.git"

	// Set the directory where you want to clone the repository
	cloneDirectory := "D:\\UpSkilling\\golang\\workdir\\webapp"

	// Clone the repository
	_, err := git.PlainClone(cloneDirectory, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})

	if err != nil && err != git.ErrRepositoryAlreadyExists {
		fmt.Println("Error cloning repository:", err)
		return
	}

	fmt.Println("Repository cloned successfully!")

	r, err := git.PlainOpen(cloneDirectory)
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return
	}

	// Get the worktree for the repository
	w, err := r.Worktree()
	if err != nil {
		fmt.Println("Error getting worktree:", err)
		return
	}

	// Edit the file
	fileName := "README.md"
	absFilePath := filepath.Join(cloneDirectory, fileName)
	file, err := os.OpenFile(absFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Write the line to add
	currentTime := time.Now()
	// Add 2 days to the current time
	twoDaysFromNow := currentTime.Add(2 * 24 * time.Hour)
	// Define the layout for the time format
	layout := "2006-01-02 15:04:05"

	// Format the time two days from now using the layout
	migrationWindow := twoDaysFromNow.Format(layout)

	projects := "group1/subgroup1/project1,   group2/subgroup2/project2"

	projectsSlice := strings.Split(projects, ",")

	for _, proj := range projectsSlice {

		line := "\n" + strings.TrimSpace(proj) + " " + migrationWindow + "\n"
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

	}

	defer file.Close()

	// Add the file to the staging area
	_, err = w.Add(fileName)
	if err != nil {
		fmt.Println("Error adding file to the staging area:", err)
		return
	}

	// Commit the changes
	commitMsg := " #2 Add a new line to the file"
	commit, err := w.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Tinesh Katta",
			Email: "tineshbabukatta@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	fmt.Println("Changes committed to", commit)

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: gitUser,
			Password: gitPass, // You can also use personal access tokens instead of passwords
		},
	})
	if err != nil {
		fmt.Println("Error pushing changes:", err)
		return
	}

	fmt.Println("Changes pushed successfully!")

}
