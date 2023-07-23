package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
	result := repoPathName("git@github.com:fatih/color.git")
	expect := "color"
	assert.Equal(t, expect, result)
}
