package main

import (
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/go-yaml/yaml"
	"log"
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
)

type config struct {
	Jobs []group `yaml:"jobs"`
}

type group struct {
	Dest  string   `yaml:"dest"`
	Repos []string `yaml:"repos"`
}

type result struct {
	lines []string
}

type outputBuilder struct {
	config *config
	result *result
}

func newOutputBuilder(config *config, result *result) *outputBuilder {
	return &outputBuilder{
		config,
		result,
	}
}

func (o *outputBuilder) writeResult() {
	for _, line := range o.result.lines {
		fmt.Println(string(line))
	}
}

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

func main() {
	config, _ := loadConfigForYaml()

	flag.Parse()
	subCommand := flag.Arg(0)

	var result result

	output := newOutputBuilder(config, &result)

	switch subCommand {
	case "install":
		for _, group := range config.Jobs {
			command := newCommandBuilder(config, output, group)
			command.moveDir()
			command.showInfo()
			command.executeCommand()
		}
	default:
		fmt.Println("Need subcommand!")
	}

	output.writeResult()
}

// 結果の構造体を作って、そこに結果を入れたい。そしてまとめた箇所で出力したい。今は出力がバラバラ
// オプションと結果を保持する構造体・メソッドを作っていく

// コマンドオプションと、ymlから読み取る設定をマージする

func loadConfigForYaml() (*config, error) {
	var configPath = flag.String("f", defaultConfigPath, "default config path")
	flag.Parse()

	f, err := os.Open(expandHomedir(*configPath))
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

func (c commandBuilder) moveDir() {
	absPath, _ := filepath.Abs(expandHomedir(c.group.Dest))
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

func expandHomedir(original string) string {
	expanded := original
	if strings.HasPrefix(original, homeDir) {
		dirname, _ := os.UserHomeDir()
		expanded = filepath.Join(dirname, original[len(homeDir):])
	}
	return expanded
}

func (c commandBuilder) executeCommand() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	for _, repo := range c.group.Repos {
		_, err := exec.Command(mainGitCommand, buildCommand(repo)...).Output()
		if err != nil {
			line := fmt.Sprintf("❌ %s \n ↪ %s", repo, err.Error())
			c.output.result.lines = append(c.output.result.lines, line)
		} else {
			c.output.result.lines = append(c.output.result.lines, fmt.Sprintf("✔ ", repo))
		}
	}
}

func buildCommand(repo string) []string {
	command := []string{subGitCommand, repo}
	return command
}
