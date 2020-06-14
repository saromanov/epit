package epit

import (
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

var (
	errNoLogger = errors.New("ExecStage: logger is not defined")
	errNoPath   = errors.New("ExecStage: path is not defined")
	errNoName   = errors.New("ExecStage: name of the stage is not defined")
	errNoStage  = errors.New("ExecStage: name of the stage is not found")
)

// ExecStage provides execution of the stage
func ExecStage(logger *zap.Logger, path, name string) error {
	if logger == nil {
		return errNoLogger
	}
	if path == "" {
		return errNoPath
	}
	if name == "" {
		return errNoName
	}
	cfg, err := loadConfig(path)
	if err != nil {
		return fmt.Errorf("ExecStage: unable to load config: %v", err)
	}
	stage, ok := cfg[name]
	if !ok {
		return errNoStage
	}
	st := Config{}
	if err := mapstructure.Decode(stage, &st); err != nil {
		return fmt.Errorf("ExecStage: unable to decode structure: %v", err)
	}

	return run(logger, name, st)
}
