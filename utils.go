package main

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

// git@github.com:fatih/color.git -> color
func repoPathName(original string) string {
	reg := regexp.MustCompile(`git@github.com:(.+?)/(.+?).git`)
	result := reg.FindAllStringSubmatch(original, -1)
	return result[0][2]
}
