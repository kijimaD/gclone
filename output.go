package main

import (
	"fmt"
)

type outputBuilder struct {
	config   *config
	result   *result
	progress *progress
}

type progress struct {
	lines []string
}

type result struct {
	lines []string
}

func newOutputBuilder(config *config, result *result, progress *progress) *outputBuilder {
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
