package hoproxy

import (
	"log"

	config "github.com/kayac/go-config"
)

// NewConfig TODO あとで書く
func NewConfig(path string) *Config {

	log.Printf("[info] loading config file: %s", path)

	cfg := &Config{}

	err := config.LoadWithEnv(&cfg, path)
	if err != nil {
		log.Fatalf("cannot load config: %s: %s", path, err)
	}

	return cfg
}
