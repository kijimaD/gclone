package main

import (
	"os"

	gclone "github.com/kijimaD/gclone/pkg"
)

func main() {
	config, _ := gclone.LoadConfigForYaml()
	output := gclone.NewOutputBuilder(config)

	for _, group := range config.Groups {
		command := gclone.NewCommandBuilder(config, output, group)
		command.PrintGroup()
	}
	output.PrintResult(os.Stdout)
}
