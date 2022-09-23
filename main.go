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
	fmt.Println(config.Repos)

	_, errExec := exec.Command(mainCommand, subCommand, "git@github.com:kijimaD/my_go.git").Output()
	if errExec != nil {
		fmt.Println("failed clone")
		fmt.Println(errExec.Error())
	} else {
		fmt.Println("successed clone")
	}
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
