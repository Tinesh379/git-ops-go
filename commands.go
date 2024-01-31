package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func GitClone() {
	// Clone the repository
	_, err := git.PlainClone(cloneDirectory, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})

	if err != nil && err != git.ErrRepositoryAlreadyExists {
		log.Fatal("Error cloning repository:", err)
		return
	}

	fmt.Println("Repository cloned successfully!")
}

func EditFile(fileName string) {

	// Edit the file
	absFilePath := filepath.Join(cloneDirectory, fileName)
	file, err := os.OpenFile(absFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}

	//add time stamp for migration
	currentTime := time.Now()
	twoDaysFromNow := currentTime.Add(2 * 24 * time.Hour)
	migrationWindow := twoDaysFromNow.Format(timeFormat)

	//projects := "group1/subgroup1/project1,   group2/subgroup2/project2"
	projects := os.Args[1]
	projectsSlice := strings.Split(projects, ",")

	// loop through the projects slugs and add projects
	for _, proj := range projectsSlice {

		line := "\n" + strings.TrimSpace(proj) + " " + migrationWindow + "\n"
		_, err = file.WriteString(line)
		if err != nil {
			log.Fatal("Error writing to file:", err)
			return
		}

	}

	defer file.Close()

}

func GitAddAndCommit(fileName string) {

	r, err := git.PlainOpen(cloneDirectory)
	if err != nil {
		log.Fatal("Error opening repository:", err)
		return
	}
	w, err := r.Worktree()
	if err != nil {
		log.Fatal("Error getting worktree:", err)
		return
	}

	// Add the file to the staging area
	_, err = w.Add(fileName)
	if err != nil {
		log.Fatal("Error adding file to the staging area:", err)
		return
	}

	// Commit the changes
	commitMsg := " #2 Add a new line to the file " + time.Now().Format(timeFormat)
	commit, err := w.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Tinesh Katta",
			Email: "tineshbabukatta@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatal("Error committing changes:", err)
		return
	}

	fmt.Println("Changes committed to", commit)
}
func GitPush() {
	r, _ := git.PlainOpen(cloneDirectory)
	err := r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: gitUser,
			Password: gitPass, // You can also use personal access tokens instead of passwords
		},
	})
	if err != nil {
		log.Fatal("Error pushing changes:", err)
		return
	}

	fmt.Println("Changes pushed successfully!")
}
