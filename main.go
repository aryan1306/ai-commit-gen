package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"ai-commits/internal"
	"ai-commits/internal/clients"

	"github.com/briandowns/spinner"
	"github.com/joho/godotenv"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

const MODEL = "gemma2:latest"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("commit-gen version %s, commit %s, built at %s\n", version, commit, date)
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	modelFlag := flag.String("model", MODEL, "The AI model to use for generating the commit message")
	flag.Parse()
	s := spinner.New(spinner.CharSets[37], 100*time.Millisecond)
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
