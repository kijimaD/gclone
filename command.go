package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

type commandBuilder struct {
	config *config
	output *outputBuilder
	group  group
}

func newCommandBuilder(config *config, output *outputBuilder, group group) *commandBuilder {
	return &commandBuilder{
		config,
		output,
		group,
	}
}

func (c commandBuilder) execute() {
	c.moveDir()
	c.showInfo()
	c.executeCommand()
}

func (c commandBuilder) moveDir() {
	absPath, _ := filepath.Abs(ExpandHomedir(c.group.Dest))
	dirErr := os.Chdir(absPath)
	if dirErr != nil {
		panic(dirErr)
	}
}

func (c commandBuilder) showInfo() {
	path, _ := os.Getwd()
	targetDir := fmt.Sprintf("Target dir: %v", path)
	reposCount := fmt.Sprintf("Repo count: %v", len(c.group.Repos))

	line := strings.Repeat("─", utf8.RuneCountInString(path))

	fmt.Println(line)
	fmt.Println(targetDir)
	fmt.Println(reposCount)
	fmt.Println(line)
}

func (c commandBuilder) executeCommand() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	for _, repo := range c.group.Repos {
		_, err := exec.Command(mainGitCommand, buildCommand(repo)...).Output()
		if err != nil {
			line := fmt.Sprintf("❌ %s \n ↪ %s", repo, err.Error())
			c.output.progress.lines = append(c.output.progress.lines, line)
		} else {
			c.output.progress.lines = append(c.output.progress.lines, fmt.Sprintf("✔ %s", repo))
		}
	}
	c.output.writeProgress()
}

func buildCommand(repo string) []string {
	command := []string{subGitCommand, repo}
	return command
}