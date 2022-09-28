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

const (
	mainGitCommand    = "git"
	subGitCommand     = "clone"
	defaultConfigPath = "config.yml"
	homeDir           = "~/"
	lineChar          = "─"
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

func (c commandBuilder) executeGroup() {
	c.prepareDir()
	c.moveDir()
	c.groupInfo()

	for _, repo := range c.group.Repos {
		c.executeCommand(repo)
	}
	c.output.appendProgress("") // newline
}

func (c commandBuilder) prepareDir() {
	// If not exist target directory, make directory
	if _, err := os.Stat(ExpandHomedir(c.group.Dest)); os.IsNotExist(err) {
		os.Mkdir(ExpandHomedir(c.group.Dest), os.ModePerm)
	}
}

func (c commandBuilder) moveDir() {
	absPath, _ := filepath.Abs(ExpandHomedir(c.group.Dest))
	dirErr := os.Chdir(absPath)
	if dirErr != nil {
		currentPath, _ := os.Getwd()
		fmt.Println("current:", currentPath)
		fmt.Println("want:", absPath)
		panic(dirErr)
	}
}

func (c commandBuilder) groupInfo() {
	currentPath, _ := os.Getwd()
	targetDir := fmt.Sprintf("Save dir: %v", currentPath)
	reposCount := fmt.Sprintf("Repo count: %v", len(c.group.Repos))
	line := strings.Repeat(lineChar, utf8.RuneCountInString(targetDir))

	c.output.appendProgress(line)
	c.output.appendProgress(targetDir)
	c.output.appendProgress(reposCount)
	c.output.appendProgress(line)
	c.output.writeProgress()
}

func (c commandBuilder) executeCommand(repo string) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	_, err := exec.Command(mainGitCommand, []string{subGitCommand, repo}...).Output()
	if err != nil {
		line := fmt.Sprintf("❌ %s \n ↪ %s", repo, err.Error())
		c.output.appendProgress(line)
		c.output.fail++
	} else {
		line := fmt.Sprintf("✔ %s", repo)
		c.output.appendProgress(line)
		c.output.success++
	}
	c.output.writeProgress()
}
