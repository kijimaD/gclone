package main

import (
	"time"

	gclone "github.com/kijimaD/gclone/pkg"
)

func main() {
	now := time.Now()
	config, _ := gclone.LoadConfigForYaml()
	var result gclone.Record
	var progress gclone.Record
	output := gclone.NewOutputBuilder(config, &result, &progress, now)

	for _, group := range config.Groups {
		command := gclone.NewCommandBuilder(config, output, group)
		command.PrintGroup()
	}
	output.PrintResult()
}
