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
