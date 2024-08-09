package config

import (
	"os"
	"path/filepath"

	"github.com/gleich/lumber/v2"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Name string `yaml:"name"`
}

func Read() Config {
	home, err := os.UserHomeDir()
	if err != nil {
		lumber.Fatal(err, "failed to get user's home directory")
	}

	path := filepath.Join(home, ".config", "tdjrl", "config.yml")
	file, err := os.ReadFile(path)
	if err != nil {
		lumber.Fatal(err, "failed to read file from", path)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		lumber.Fatal(err, "failed to unmarshal yaml file with content:", string(file))
	}
	return config
}
