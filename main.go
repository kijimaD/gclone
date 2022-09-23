package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"os/exec"
	"time"
)

const (
	mainCommand = "git"
	subCommand  = "clone"
)

func main() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	fmt.Println("start cloning...", time.Now())
	_, err := exec.Command(mainCommand, subCommand, "git@github.com:kijimaD/my_go.git").Output()
	if err != nil {
		fmt.Println("failed clone")
		fmt.Println(err.Error())
	} else {
		fmt.Println("successed clone")
	}
	s.Stop()
}
