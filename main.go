package main

import (
	"flag"
	"fmt"
)

const (
	mainGitCommand    = "git"
	subGitCommand     = "clone"
	defaultConfigPath = "config.yml"
	homeDir           = "~/"
)

type config struct {
	Groups []group `yaml:"groups"`
}

type group struct {
	Dest  string   `yaml:"dest"`
	Repos []string `yaml:"repos"`
}

type outputBuilder struct {
	config *config
	result *result
}

type result struct {
	lines []string
}

func newOutputBuilder(config *config, result *result) *outputBuilder {
	return &outputBuilder{
		config,
		result,
	}
}

func (o *outputBuilder) writeResult() {
	for _, line := range o.result.lines {
		fmt.Println(string(line))
	}
}

func main() {
	config, _ := LoadConfigForYaml()
	flag.Parse()
	subCommand := flag.Arg(0)

	var result result

	output := newOutputBuilder(config, &result)

	switch subCommand {
	case "install":
		for _, group := range config.Groups {
			command := newCommandBuilder(config, output, group)
			command.execute()
		}
	default:
		fmt.Println("Need subcommand!")
	}

	output.writeResult()
}

// 結果の構造体を作って、そこに結果を入れたい。そしてまとめた箇所で出力したい。今は出力がバラバラ
// オプションと結果を保持する構造体・メソッドを作っていく
// コマンドオプションと、ymlから読み取る設定をマージする
