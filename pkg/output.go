package gclone

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// グループ・コマンド全体を通して使用される
type outputBuilder struct {
	config   *config
	result   *Record // 最後に1回だけ実行される類いの表示内容
	progress *Record // 途中で何回か実行されそのたびに内容がリセットされる類の表示内容
	success  int
	fail     int
	now      time.Time
	writer   io.Writer
	flash    string
}

type Record struct {
	lines []string
}

func NewOutputBuilder(config *config, w io.Writer) *outputBuilder {
	record := Record{}
	progress := Record{}
	now := time.Now()
	return &outputBuilder{
		config,
		&record,
		&progress,
		0,
		0,
		now,
		w,
		"",
	}
}

func (o *outputBuilder) printProgress() {
	for _, line := range o.progress.lines {
		fmt.Fprintln(o.writer, string(line))
	}
	o.progress.lines = []string{} // flush
}

func (o *outputBuilder) printNewLine() {
	fmt.Fprintln(o.writer, "")
}

func (o *outputBuilder) printFlash() {
	fmt.Fprintf(o.writer, "%s\r", string(o.flash))
}

func (o *outputBuilder) PrintResult() {
	tmpl := `
done!
Success: %d
Fail: %d
Process: %vms
`
	fmt.Fprintf(o.writer, tmpl, o.success, o.fail, time.Since(o.now).Milliseconds())
	for _, line := range o.result.lines {
		fmt.Fprintln(o.writer, string(line))
	}
}

func (o *outputBuilder) appendProgress(line string) {
	o.progress.lines = append(o.progress.lines, line)
}

func (o *outputBuilder) appendFlash(str string) {
	slice := []string{o.flash, str}
	o.flash = strings.Join(slice, "")
}

func (o *outputBuilder) appendResult(line string) {
	o.result.lines = append(o.result.lines, line)
}

func (o *outputBuilder) flushFlash() {
	o.flash = ""
}
