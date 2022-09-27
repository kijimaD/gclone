package main

import (
	"testing"
)

// func TestMoveDir(t *testing.T) {
// 	{
// 		moveDir("~/")

// 		home, _ := os.UserHomeDir()
// 		pwd, _ := os.Getwd()

// 		if !(pwd == home) {
// 			t.Error(pwd, home)
// 		}
// 	}

// 	{
// 		moveDir("/")

// 		root := "/"
// 		pwd, _ := os.Getwd()

// 		if !(pwd == root) {
// 			t.Error(pwd, root)
// 		}
// 	}
// }

func TestBuildCommand(t *testing.T) {
	repo := "git@github.com:kijimaD/gclone.git"

	result := buildCommand(repo)
	expect := []string{subGitCommand, repo}

	if !(result[0] == expect[0]) || !(result[1] == expect[1]) {
		t.Error("result: ", result, "expect:", expect)
	}
}
