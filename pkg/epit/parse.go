package epit

import (
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/saromanov/cowrow"
)

type Stage struct {
	Command string
}

// Config provides definition of configuration
type Config map[string]interface{}

// loadConfig provides loading of configuration file
func loadConfig(path string) (Config, error) {

	if path == "" {
		return nil, errors.New("path to config is empty")
	}

	cfg := Config{}
	err := cowrow.LoadByPath(path, &cfg)
	if err != nil {
		return nil, err
	}

	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("unable to validate config: %v", err)
	}

	return cfg, nil
}

func validateConfig(cfg Config) error {
	if len(cfg) == 0 {
		return errors.New("stage on the config is not defined")
	}

	return nil
}
