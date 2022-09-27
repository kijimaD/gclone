package main

import ()

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

func main() {
	config, _ := LoadConfigForYaml()
	var result result
	var progress progress
	output := newOutputBuilder(config, &result, &progress)

	for _, group := range config.Groups {
		command := newCommandBuilder(config, output, group)
		command.execute()
	}
	output.writeResult()
}

// 結果の構造体を作って、そこに結果を入れたい。そしてまとめた箇所で出力したい。今は出力がバラバラ
// オプションと結果を保持する構造体・メソッドを作っていく
// コマンドオプションと、ymlから読み取る設定をマージする
