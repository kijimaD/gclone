package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/go-yaml/yaml"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	mainCommand = "git"
	subCommand  = "clone"
)

type config struct {
	Repos []string `yaml:"ident"`
}

func main() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	fmt.Println("start cloning...", time.Now())

	config, _ := loadConfigForYaml()

	for _, repo := range config.Repos {
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
	command := []string {subCommand, repo}
	return command
}

func loadConfigForYaml() (*config, error) {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

// sshダウンロードが使えるかチェック
// https/sshを選べるようにする
// ディレクトリが存在したらスキップ
// ダウンロード先ディレクトリを設定できるようにする
// 設定ファイルの雛形を作成できるようにする
