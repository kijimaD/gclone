package main

import (
	"fmt"
)

type outputBuilder struct {
	config   *config
	result   *record
	progress *record
	success  int
	fail     int
}

type record struct {
	lines []string
}

func newOutputBuilder(config *config, result *record, progress *record) *outputBuilder {
	return &outputBuilder{
		config,
		result,
		progress,
		0,
		0,
	}
}

func (o *outputBuilder) writeProgress() {
	for _, line := range o.progress.lines {
		fmt.Println(string(line))
	}
	o.progress.lines = []string{""}
}

func (o *outputBuilder) writeResult() {
	fmt.Println("done")
	fmt.Println("Success:", o.success)
	fmt.Println("Fail:", o.fail)
	for _, line := range o.result.lines {
		fmt.Println(string(line))
	}
}

func (o *outputBuilder) appendProgress(line string) {
	o.progress.lines = append(o.progress.lines, line)
}

func (o *outputBuilder) appendResult(line string) {
	o.result.lines = append(o.result.lines, line)
}
