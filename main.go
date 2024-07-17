package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		log.Fatal("No $HOME env")
	}

	targetDir := filepath.Join(homeDir, "Repos", "obsidian-notes") // One repo for now

	if err := os.Chdir(targetDir); err != nil {
		log.Fatal("Error changing directory:", err)
	}

	fmt.Println("Current Directory: ", targetDir)

	cmd := exec.Command("git", "add", ".")
	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	cmd = exec.Command("git", "commit", "-m", currentTime+" - Auto pushed")
	stdout, err := cmd.Output()
	fmt.Println(string(stdout))
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "push")
	stdout, err = cmd.Output()
	fmt.Println(string(stdout))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("======================")
	fmt.Println("Successfully pushed the repos")
}
