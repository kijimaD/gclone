package gclone

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintProgress(t *testing.T) {
	config := config{}
	buf := bytes.Buffer{}
	o := NewOutputBuilder(&config, &buf)
	o.progress.lines = []string{"hello", "world"}

	o.printProgress()

	actual := buf.String()
	expect := "hello\nworld\n"

	assert.Equal(t, expect, actual)

	// 1度表示すると、内部のprogressはリセットされる
	bufafter := bytes.Buffer{}
	o.printProgress()
	actual = bufafter.String()
	expect = ""
	assert.Equal(t, expect, actual)
}

func TestPrintResult(t *testing.T) {
	config := config{}
	buf := bytes.Buffer{}
	o := NewOutputBuilder(&config, &buf)
	o.result.lines = []string{"hello", "world"}

	o.PrintResult()

	actual := buf.String()
	expect := `
done!
Success: 0
Fail: 0
Process: 0ms
hello
world
`
	assert.Equal(t, expect, actual)
}
