package gclone

import (
	"fmt"
	"time"
)

type outputBuilder struct {
	config   *config
	result   *Record
	progress *Record
	success  int
	fail     int
	now      time.Time
}

type Record struct {
	lines []string
}

func NewOutputBuilder(config *config, result *Record, progress *Record, now time.Time) *outputBuilder {
	return &outputBuilder{
		config,
		result,
		progress,
		0,
		0,
		now,
	}
}

func (o *outputBuilder) printProgress() {
	for _, line := range o.progress.lines {
		fmt.Println(string(line))
	}
	o.progress.lines = []string{}
}

func (o *outputBuilder) PrintResult() {
	fmt.Printf("\ndone!\n")
	fmt.Println("Success:", o.success)
	fmt.Println("Fail:", o.fail)
	fmt.Printf("Process: %vms\n", time.Since(o.now).Milliseconds())
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
