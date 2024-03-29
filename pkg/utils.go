package gclone

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-yaml/yaml"
)

func LoadConfigForYaml() (*config, error) {
	var configPath = flag.String("f", defaultConfigPath, "default config path")
	flag.Parse()

	f, err := os.Open(ExpandHomedir(*configPath))
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

func ExpandHomedir(original string) string {
	expanded := original
	if strings.HasPrefix(original, homeDir) {
		dirname, _ := os.UserHomeDir()
		expanded = filepath.Join(dirname, original[len(homeDir):])
	}
	return expanded
}

// (example) "git@github.com:fatih/color.git" -> "color"
func repoPathName(original string) string {
	regMatch := regexp.MustCompile(`(https://|git@)github.com(:|/)(.+?)/(.+?)$`)
	match := regMatch.FindAllStringSubmatch(original, -1)

	regReplace := regexp.MustCompile(`.git`) // ついているときもあるので除去する
	result := match[0][4]
	result = regReplace.ReplaceAllString(result, "")
	return result
}
