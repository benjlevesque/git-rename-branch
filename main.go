package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func getCurrentBranch() string {
	output := new(strings.Builder)

	c := exec.Command("git", "branch", "--show-current")
	c.Stdout = output
	c.Stderr = output
	if err := c.Run(); err != nil {
		log.Println(output)
		log.Fatal(err)
	}
	return strings.TrimRight(output.String(), "\n")

}

func editValue(current string) string {
	f, err := os.CreateTemp("", "branch*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())

	if _, err := f.WriteString(current); err != nil {
		log.Fatal(err)
	}
	editor := os.Getenv("EDITOR")
	output := new(strings.Builder)
	cmd := exec.Command(editor, f.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		log.Println(output)
		log.Fatal(err)
	}
	raw, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	return string(raw)
}

func renameBranch(newName string) {
	output := new(strings.Builder)

	c := exec.Command("git", "branch", "--move", newName)
	c.Stdout = output
	c.Stderr = output
	if err := c.Run(); err != nil {
		log.Println(output)
		log.Fatal(err)
	}
}

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionFlag = flag.Bool("version", false, "Display version")

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Printf("Version %s, commit %s, built at %s\n", version, commit, date)
		return
	}

	branchName := getCurrentBranch()
	newBranchName := strings.Split(editValue(branchName), "\n")[0]
	renameBranch(newBranchName)
}
