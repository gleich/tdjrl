package config

import (
	"os"

	"github.com/gleich/lumber/v2"
	"gopkg.in/yaml.v3"
)

const FILENAME = "test_config.yml"

type Config struct {
	Name string `yaml:"name"`
}

func Load() Config {
	file, err := os.ReadFile(FILENAME)
	if err != nil {
		lumber.FatalMsg("Failed to read from file: ", FILENAME)
	}

	var loadedConfig Config
	err = yaml.Unmarshal(file, &loadedConfig)
	if err != nil {
		lumber.FatalMsg("Failed to unmarshal yaml file")
	}

	return loadedConfig
}
