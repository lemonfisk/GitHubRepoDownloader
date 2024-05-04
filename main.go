package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Repo struct {
	Name string
	Link string
}

func cloneRepositories(cloneDir string, repoName string, repoLink string) {
	repoPath := filepath.Join(cloneDir, repoName)
	cmd := exec.Command("git", "clone", repoLink, repoPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Repository clonning error %s: %v\n", repoName, err)
	} else {
		fmt.Printf("Repo %s download successful \n", repoName)
	}
}

func main() {

	dataJson := "configs/repo.json"
	cloneDir := "downloads"

	data, err := os.ReadFile(dataJson)

	if err != nil {
		log.Fatal(err)
	}

	var repositories []Repo

	err = json.Unmarshal(data, &repositories)

	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repositories {
		cloneRepositories(cloneDir, repo.Name, repo.Link)
	}
}
