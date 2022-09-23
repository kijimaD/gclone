package main

import (
	"fmt"
	"os/exec"
	"time"
	// "time"
	// "github.com/briandowns/spinner"
)

func main() {
	// s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	// s.Start()
	fmt.Println("Start...", time.Now())
	out, err := exec.Command("git", "clone", "git@github.com:kijimaD/my_go.git").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(out))
	// s.Stop()
}
