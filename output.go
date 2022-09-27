package main

import (
	"fmt"
)

type outputBuilder struct {
	config   *config
	result   *record
	progress *record
}

type record struct {
	lines []string
}

func newOutputBuilder(config *config, result *record, progress *record) *outputBuilder {
	return &outputBuilder{
		config,
		result,
		progress,
	}
}

func (o *outputBuilder) writeProgress() {
	for _, line := range o.progress.lines {
		fmt.Println(string(line))
	}
	o.progress.lines = []string{""}
}

func (o *outputBuilder) writeResult() {
	for _, line := range o.result.lines {
		fmt.Println(string(line))
	}
}
