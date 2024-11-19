package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func StageAllFiles(r *bufio.Reader) {
	fmt.Printf("Stage all uncommited files? (y/n) ")

	confirmStage, err := r.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading input")
		os.Exit(1)
	}

	if strings.ToLower(strings.TrimSpace(confirmStage)) == "y" {
		exec.Command("git", "add", ".").Run()
		fmt.Println("Files staged")
	} else {
		fmt.Println("Exiting...")
		os.Exit(1)
	}
}

func GenerateDiff() []byte {
	diffCmd := exec.Command("git", "diff", "--staged")
	diffOut, err := diffCmd.Output()
	if err != nil {
		fmt.Println("Error running git diff")
		os.Exit(1)
	}
	return diffOut
}