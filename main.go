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

	fmt.Println("start cloning: ", time.Now().Format("2006-01-02 15:04:05"))

	config, _ := loadConfigForYaml()
	moveDir(config.Dest)
	currentDir()
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

func currentDir() {
	p, _ := os.Getwd()
	fmt.Println("current dir: ", p)
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
