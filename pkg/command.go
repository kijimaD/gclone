package gclone

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	mainGitCommand     = "git"
	subGitCommand      = "clone"
	defaultConfigPath  = "config.yml"
	homeDir            = "~/"
	horizontalLineChar = "─"
	progressChar       = "."
)

// 1つ1つのシェルコマンドを指す構造体
type commandBuilder struct {
	config *config
	output *outputBuilder
	group  group
}

func NewCommandBuilder(config *config, output *outputBuilder, group group) *commandBuilder {
	return &commandBuilder{
		config,
		output,
		group,
	}
}

func (c commandBuilder) ExecGroup() {
	c.prepareDir()
	c.moveDir()
	c.writeGroupInfo()

	for _, repo := range c.group.Repos {
		c.executeCommand(repo)
	}
	c.output.appendProgress("") // newline
}

func (c commandBuilder) prepareDir() {
	// If not exist target directory, make directory
	if _, err := os.Stat(ExpandHomedir(c.group.Dest)); os.IsNotExist(err) {
		os.Mkdir(ExpandHomedir(c.group.Dest), os.ModePerm)
	}
}

func (c commandBuilder) moveDir() {
	absPath, _ := filepath.Abs(ExpandHomedir(c.group.Dest))
	dirErr := os.Chdir(absPath)
	if dirErr != nil {
		currentPath, _ := os.Getwd()
		log.Fatalf("current: %s, want: %s", currentPath, absPath)
	}
}

func (c commandBuilder) writeGroupInfo() {
	currentPath, _ := os.Getwd()
	targetDir := fmt.Sprintf("Save dir: %v", currentPath)
	reposCount := fmt.Sprintf("Repo count: %v", len(c.group.Repos))
	line := strings.Repeat(horizontalLineChar, utf8.RuneCountInString(targetDir))

	c.output.appendProgress(line)
	c.output.appendProgress(targetDir)
	c.output.appendProgress(reposCount)
	c.output.appendProgress(line)
	c.output.printProgress()
}

// 実行中にプログレスバーを表示するために非同期実行にしている
// FIXME: 汚いのでリファクタ
func (c commandBuilder) executeCommand(repo string) {
	type output struct {
		out []byte
		err error
	}

	now := time.Now()
	ch := make(chan output)
	go func() {
		out, err := exec.Command(mainGitCommand, []string{subGitCommand, repo}...).CombinedOutput()
		ch <- output{out, err}
	}()

	fmt.Print(fmt.Sprintf("%s\n", repo))
progress:
	for {
		select {
		case output := <-ch:
			if output.err != nil {
				line := fmt.Sprintf(" ❌ \n ↪ %s \n ↪ %s", string(output.err.Error()), string(output.out))
				c.output.appendProgress(line)
				c.output.fail++

				// 失敗した場合にgit pullを実行する
				str := repoPathName(repo)

				// ディレクトリが存在しない場合は処理を中断
				_, err := os.Stat(str)
				if err != nil {
					panic(err)
				}
				c.moveDir()
				currentPath, _ := os.Getwd()
				dirErr := os.Chdir(path.Join(currentPath, str))
				if dirErr != nil {
					panic(dirErr)
				}
				// TODO: git pullもチャンネル化する
				out, _ := exec.Command(mainGitCommand, []string{"pull"}...).CombinedOutput()
				c.output.appendProgress(string(out))
			} else {
				time := fmt.Sprintf("%ds", int(time.Since(now).Seconds()))
				line := fmt.Sprintf(" ✔ (%s)", time)
				c.output.appendProgress(line)
				c.output.success++
			}
			c.output.flushFlash()
			break progress
		default:
			c.output.appendFlash(progressChar)
			c.output.printFlash()
			timer := time.NewTimer(1 * time.Second)
			<-timer.C
		}
	}

	c.output.printProgress()
}
