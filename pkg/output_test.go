package gclone

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintProgress(t *testing.T) {
	config := config{}
	o := NewOutputBuilder(&config)
	o.progress.lines = []string{"hello", "world"}

	buf := bytes.Buffer{}
	o.printProgress(&buf)

	actual := buf.String()
	expect := "hello\nworld\n"

	assert.Equal(t, expect, actual)

	// 1度表示すると、内部のprogressはリセットされる
	bufafter := bytes.Buffer{}
	o.printProgress(&bufafter)
	actual = bufafter.String()
	expect = ""
	assert.Equal(t, expect, actual)
}
