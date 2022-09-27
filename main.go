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
	mainGitCommand = "git"
	subGitCommand  = "clone"
	defaultConfigPath  = "config.yml"
	homeDir        = "~/"
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
	return &outputBuilder {
		config,
		result,
	}
}

func (o *outputBuilder) writeResult() {
	for _, line := range o.result.lines {
		fmt.Println(string(line))
	}
}

func main() {
	config, _ := loadConfigForYaml()

	flag.Parse()
	subCommand := flag.Arg(0)

	var result result
	result.lines = make([]string, 2)

	switch subCommand {
	case "install":
		for _, c := range config.Jobs {
			moveDir(c.Dest)
			showInfo(c, &result)
			executeCommand(c.Repos, &result)
		}
	default:
		fmt.Println("Need subcommand!")
	}

	builder := newOutputBuilder(config, &result)
	builder.writeResult()
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

func moveDir(path string) {
	absPath, _ := filepath.Abs(expandHomedir(path))
	dirErr := os.Chdir(absPath)
	if dirErr != nil {
		panic(dirErr)
	}
}

func showInfo(group group, result *result) {
	path, _ := os.Getwd()
	targetDir := fmt.Sprintf("Target dir: %v", path)
	reposCount := fmt.Sprintf("Repo count: %v", len(group.Repos))

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

func executeCommand(repos []string, result *result) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	for _, repo := range repos {
		_, err := exec.Command(mainGitCommand, buildCommand(repo)...).Output()
		if err != nil {
			line := fmt.Sprintf("❌ %s \n ↪ %s", repo, err.Error())
			result.lines = append(result.lines, line)
		} else {
			result.lines = append(result.lines, fmt.Sprintf("✔ ", repo))
		}
	}
}

func buildCommand(repo string) []string {
	command := []string{subGitCommand, repo}
	return command
}
