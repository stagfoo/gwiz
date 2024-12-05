package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

// CommitOptions defines the structure for commit types from YAML
type CommitOptions struct {
	CommitTypes []CommitType `yaml:"commit_types"`
}

// CommitType represents a single commit type option
type CommitType struct {
	ID     string `yaml:"id"`
	Emoji  string `yaml:"emoji"`
	Prefix string `yaml:"prefix"`
	Name   string `yaml:"name"`
}

func main() {
	commitOptions, err := loadCommitOptions()
	if err != nil {
		fmt.Println("Error loading commit options:", err)
		return
	}
	gwiz(commitOptions)
}

func loadCommitOptions() (CommitOptions, error) {
	var options CommitOptions
	
	// Try to read from config file in multiple locations
	configPaths := []string{
		"./commit-types.yaml",
		"~/.config/gwiz/commit-types.yaml",
		"/etc/gwiz/commit-types.yaml",
	}
	
	for _, path := range configPaths {
		data, err := os.ReadFile(path)
		if err == nil {
			err = yaml.Unmarshal(data, &options)
			if err != nil {
				return CommitOptions{}, fmt.Errorf("error parsing YAML from %s: %v", path, err)
			}
			return options, nil
		}
	}
	
	// If no config file found, return default options
	return defaultCommitOptions(), nil
}

func defaultCommitOptions() CommitOptions {
	return CommitOptions{
		CommitTypes: []CommitType{
			{ID: "1", Emoji: "✨", Prefix: "feature:", Name: "feature"},
			{ID: "2", Emoji: "🍱", Prefix: "refactor:", Name: "refactor"},
			{ID: "3", Emoji: "🧼", Prefix: "chore:", Name: "chore"},
			{ID: "4", Emoji: "💿", Prefix: "asset:", Name: "asset"},
			{ID: "5", Emoji: "🐞", Prefix: "fix:", Name: "fix"},
			{ID: "6", Emoji: "🚀", Prefix: "release:", Name: "release"},
			{ID: "7", Emoji: "📚", Prefix: "docs:", Name: "docs"},
			{ID: "8", Emoji: "🤖", Prefix: "test:", Name: "test"},
			{ID: "9", Emoji: "🚓", Prefix: "security:", Name: "security"},
			{ID: "10", Emoji: "↩️", Prefix: "revert:", Name: "revert"},
		},
	}
}

func gwiz(options CommitOptions) {
	reader := bufio.NewReader(os.Stdin)
	
	// Print commit type options
	fmt.Println("Select the type of commit:")
	for _, commitType := range options.CommitTypes {
		fmt.Printf("%s) %s %s\n", commitType.ID, commitType.Emoji, commitType.Name)
	}
	fmt.Println("0) Cancel")
	
	commitTypeInput, _ := reader.ReadString('\n')
	commitTypeInput = strings.TrimSpace(commitTypeInput)
	
	// Find selected commit type
	var selectedCommitType *CommitType
	for _, commitType := range options.CommitTypes {
		if commitType.ID == commitTypeInput {
			selectedCommitType = &commitType
			break
		}
	}
	
	if commitTypeInput == "0" {
		fmt.Println("Commit canceled.")
		return
	}
	
	if selectedCommitType == nil {
		fmt.Println("Invalid option.")
		return
	}
	
	// Prompt for commit message
	fmt.Print("Enter the commit message: ")
	commitMessage, _ := reader.ReadString('\n')
	commitMessage = strings.TrimSpace(commitMessage)
	
	if commitMessage == "" {
		fmt.Println("Commit message cannot be empty.")
		return
	}
	
	// Construct full commit message
	fullMessage := fmt.Sprintf("%s %s %s", selectedCommitType.Emoji, selectedCommitType.Prefix, commitMessage)
	
	// Prompt for hook skipping
	fmt.Print("Do you want to skip the git hooks? (y/n): ")
	skipHooks, _ := reader.ReadString('\n')
	skipHooks = strings.TrimSpace(skipHooks)
	
	// Prepare commit command
	cmd := fmt.Sprintf("git commit -m \"%s\"", fullMessage)
	if strings.ToLower(skipHooks) == "y" {
		cmd += " -n"
	}
	
	fmt.Println("Executing:", cmd)
	err := executeCommand(cmd)
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}

func executeCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}