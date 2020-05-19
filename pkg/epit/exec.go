package epit

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

// ExecStage provides execution of the stage
func ExecStage(logger *zap.Logger, path, name string) error {
	cfg, err := loadConfig(path)
	if err != nil {
		return fmt.Errorf("ExecStage: unable to load config: %v", err)
	}
	stage, ok := cfg[name]
	if !ok {
		return fmt.Errorf("name of the stage is not found")
	}
	st := Config{}
	if err := mapstructure.Decode(stage, &st); err != nil {
		return fmt.Errorf("unable to decode structure: %v", err)
	}

	return run(name, st)
}
