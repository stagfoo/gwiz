package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	gwiz()
}

func gwiz() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select the type of commit:")
	fmt.Println("1) ✨ feature")
	fmt.Println("2) 🍱 refactor")
	fmt.Println("3) 🧼 chore")
	fmt.Println("4) 🐞 fix")
	fmt.Println("5) 🚀 release")
	fmt.Println("6) 📚 docs")
	fmt.Println("7) 🤖 test")
	fmt.Println("8) 🚓 security")
	fmt.Println("9) ↩️ revert")
	fmt.Println("0) Cancel")

	commitType, _ := reader.ReadString('\n')
	commitType = strings.TrimSpace(commitType)
	var commitPrefix string

	switch commitType {
	case "1":
		commitPrefix = "✨ feature:"
	case "2":
		commitPrefix = "🍱 refactor:"
	case "3":
		commitPrefix = "🧼 chore:"
	case "4":
		commitPrefix = "🐞 fix:"
	case "5":
		commitPrefix = "🚀 release:"
	case "6":
		commitPrefix = "📚 docs:"
	case "7":
		commitPrefix = "🤖 test:"
	case "8":
		commitPrefix = "🚓 security:"
	case "9":
		commitPrefix = "↩️ revert:"
	case "0":
		fmt.Println("Commit canceled.")
		return
	default:
		fmt.Println("Invalid option.")
		return
	}

	fmt.Print("Enter the commit message: ")
	commitMessage, _ := reader.ReadString('\n')
	commitMessage = strings.TrimSpace(commitMessage)

	if commitMessage == "" {
		fmt.Println("Commit message cannot be empty.")
		return
	}

	fullMessage := fmt.Sprintf("%s %s", commitPrefix, commitMessage)

	fmt.Print("Do you want to skip the git hooks? (y/n): ")
	skipHooks, _ := reader.ReadString('\n')
	skipHooks = strings.TrimSpace(skipHooks)

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

