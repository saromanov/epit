package epit

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

var (
	errNoLogger = errors.New("ExecStage: logger is not defined")
	errNoPath   = errors.New("ExecStage: path is not defined")
	errNoName   = errors.New("ExecStage: name of the stage is not defined")
	errNoStage  = errors.New("ExecStage: name of the stage is not found")
)

// Exec provides executing of the stages
func Exec(logger *zap.Logger, path, pattern string) error {
	if err := validate(logger, path); err != nil {
		return fmt.Errorf("unable to validate input: %v", err)
	}
	cfg, err := loadConfig(path)
	if err != nil {
		return fmt.Errorf("ExecStage: unable to load config: %v", err)
	}
	for name, param := range cfg {
		matched, err := regexp.MatchString(pattern, name)
		if err != nil {
			logger.Error("unable to match stage pattern", zap.Error(err))
			continue
		}
		if !matched {
			continue
		}
		if err := execInner(logger, name, param); err != nil {
			logger.Error("unable to execute stage", zap.Error(err))
			continue
		}
	}
	return nil

}

// execInner provides execution of the stage
func execInner(logger *zap.Logger, name string, stage interface{}) error {
	st := Config{}
	if err := mapstructure.Decode(stage, &st); err != nil {
		return fmt.Errorf("ExecStage: unable to decode structure: %v", err)
	}

	return run(logger, name, st)
}

func validate(logger *zap.Logger, path string) error {
	if logger == nil {
		return errNoLogger
	}
	if path == "" {
		return errNoPath
	}
	return nil
}
