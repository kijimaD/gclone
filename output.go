package main

import (
	"fmt"
)

type outputBuilder struct {
	config *config
	result *result
}

type result struct {
	lines []string
}

func newOutputBuilder(config *config, result *result) *outputBuilder {
	return &outputBuilder{
		config,
		result,
	}
}

func (o *outputBuilder) writeResult() {
	for _, line := range o.result.lines {
		fmt.Println(string(line))
	}
}
