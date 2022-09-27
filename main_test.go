package main

import (
	"os"
	"testing"
)

func TestMoveDir(t *testing.T) {
	{
		moveDir("~/")

		home, _ := os.UserHomeDir()
		pwd, _ := os.Getwd()

		if !(pwd == home) {
			t.Error(pwd, home)
		}
	}

	{
		moveDir("/")

		root := "/"
		pwd, _ := os.Getwd()

		if !(pwd == root) {
			t.Error(pwd, root)
		}
	}
}

func TestExpandHomeDir(t *testing.T) {
	// expand ~
	{
		home, _ := os.UserHomeDir()

		result := expandHomedir("~/")
		expect := home
		if !(result == expect) {
			t.Error("result: ", result, "expect:", expect)
		}
	}

	// not expand
	{
		result := expandHomedir("/home/user")
		expect := "/home/user"
		if !(result == expect) {
			t.Error("result: ", result, "expect:", expect)
		}
	}
}

func TestBuildCommand(t *testing.T) {
	repo := "git@github.com:kijimaD/gclone.git"

	result := buildCommand(repo)
	expect := []string{subGitCommand, repo}

	if !(result[0] == expect[0]) || !(result[1] == expect[1]) {
		t.Error("result: ", result, "expect:", expect)
	}
}
