package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "git",
		Short: "Automate Git commands with GitHub integration",
		Run: func(cmd *cobra.Command, args []string) {
			createGitHubRepositoryAndPushChanges()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createGitHubRepositoryAndPushChanges() {
	repoName := getCurrentDirectoryName()
	fmt.Println("Directory Name:", repoName)
	repoURL := createGitHubRepository(repoName)
	initializeGit()
	addRemoteOrigin(repoURL)
	gitCommand("add", ".")
	commitMsg := getUserInput("Enter commit message: ")
	gitCommand("commit", "-m", commitMsg)
	gitCommand("branch", "-M", "main")
	gitCommand("push", "-u", "origin", "main")
}

func getCurrentDirectoryName() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}
	dirParts := strings.Split(dir, string(os.PathSeparator))
	return dirParts[len(dirParts)-1]
}

func createGitHubRepository(repoName string) string {
	fmt.Println("Repository Name:", repoName)
	cmd := exec.Command("gh", "repo", "create", repoName, "--public")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error creating GitHub repository:", err)
		os.Exit(1)
	}
	repoURL := strings.Split(string(output), "\n")[0]
	return repoURL
}
func initializeGit() {
	gitCommand("init")
}

func addRemoteOrigin(url string) {
	gitCommand("remote", "add", "origin", url)
}

func gitCommand(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing git command:", strings.Join(args, " "))
		os.Exit(1)
	}
}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
