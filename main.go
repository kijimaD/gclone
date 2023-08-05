package main

import (
	"time"
)

type config struct {
	Groups []group `yaml:"groups"`
}

type group struct {
	Dest  string   `yaml:"dest"`
	Repos []string `yaml:"repos"`
}

func main() {
	now := time.Now()
	config, _ := LoadConfigForYaml()
	var result record
	var progress record
	output := newOutputBuilder(config, &result, &progress, now)

	for _, group := range config.Groups {
		command := newCommandBuilder(config, output, group)
		command.printGroup()
	}
	output.writeResult()
}
