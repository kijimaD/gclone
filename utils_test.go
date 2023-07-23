package main

import (
	"os"
	"testing"
)

func TestExpandHomeDir(t *testing.T) {
	// expand ~
	{
		home, _ := os.UserHomeDir()

		result := ExpandHomedir("~/")
		expect := home
		if !(result == expect) {
			t.Error("result: ", result, "expect:", expect)
		}
	}

	// not expand
	{
		result := ExpandHomedir("/home/user")
		expect := "/home/user"
		if !(result == expect) {
			t.Error("result: ", result, "expect:", expect)
		}
	}
}

func TestRepoPathName(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "1",
			input:  "git@github.com:fatih/color.git",
			expect: "color",
		},
		{
			name:   "2",
			input:  "git@github.com:kijimaD/cloner",
			expect: "cloner",
		},
		{
			name:   "3",
			input:  "https://github.com/kijimaD/gclone.git",
			expect: "gclone",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repoPathName(tt.input)
			if got != tt.expect {
				t.Errorf("got %s want %s", got, tt.expect)
			}
		})
	}
}
