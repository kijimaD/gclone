package main

import (
	"flag"
	"github.com/go-yaml/yaml"
	"log"
	"os"
	"path/filepath"
	"strings"
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
