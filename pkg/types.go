package gclone

type config struct {
	Groups []group `yaml:"groups"`
}

type group struct {
	Dest  string   `yaml:"dest"`
	Repos []string `yaml:"repos"`
}
