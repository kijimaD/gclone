package main

import (
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
	mainCommand = "git"
	subCommand  = "clone"
)

type config struct {
	Dest  string   `yaml:"dest"`
	Repos []string `yaml:"repos"`
}

func main() {
	// spinner
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	config, _ := loadConfigForYaml()
	moveDir(config.Dest)
	showInfo(config)
	executeCommand(config.Repos)
}

func loadConfigForYaml() (*config, error) {
	f, err := os.Open("config.yml")
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

func showInfo(config *config) {
	path, _ := os.Getwd()
	targetDir := fmt.Sprintf("Target dir: %v", path)
	reposCount := fmt.Sprintf("Repo count: %v", len(config.Repos))

	line := strings.Repeat("─", utf8.RuneCountInString(path))

	fmt.Println(line)
	fmt.Println(targetDir)
	fmt.Println(reposCount)
	fmt.Println(line)
}

func expandHomedir(path string) string {
	var expanded string
	if strings.HasPrefix(path, "~/") {
		dirname, _ := os.UserHomeDir()
		expanded = filepath.Join(dirname, path[2:])
	}
	return expanded
}

func executeCommand(repos []string) {
	for _, repo := range repos {
		_, err := exec.Command(mainCommand, buildCommand(repo)...).Output()
		if err != nil {
			fmt.Println("❌ ", repo)
			fmt.Println(" ↪", err.Error())
		} else {
			fmt.Println("✔ ", repo)
		}
	}
}

func buildCommand(repo string) []string {
	command := []string{subCommand, repo}
	return command
}

// ディレクトリが存在したらスキップ
// 設定ファイルの雛形を作成できるようにする
