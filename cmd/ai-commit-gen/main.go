package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aryan1306/ai-commit-gen/internal"
	"github.com/aryan1306/ai-commit-gen/internal/clients"

	"github.com/briandowns/spinner"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	versionFlag := flag.Bool("version", false, "Print version information")
	modelFlag := flag.String("model", "", "The AI model to use for generating the commit message")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("ai-commit-gen version %s, commit %s, built at %s\n", version, commit, date)
		os.Exit(0)
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " ✨Generating commit message..."
	s.Color("green")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose an AI model:")
	fmt.Println("1. OpenAI")
	fmt.Println("2. Ollama")
	modelChoice, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	
	internal.StageAllFiles(reader)

	fmt.Println("generate commit message?")
	generateCommitMessage, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	if strings.TrimSpace(generateCommitMessage) == "y" {
		switch strings.TrimSpace(modelChoice) {
		case "1":
			clients.OpenAiClient(s)
		case "2":
			clients.OllamaClient(s, *modelFlag)
		default:
			fmt.Println("Invalid choice")
		}
	}

}
