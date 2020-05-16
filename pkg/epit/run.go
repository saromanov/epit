package epit

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// run provides running of the stage
func run(cfg Config) error {
	envs, ok := cfg[env]
	if ok {
		addEnvVariables(envs.([]interface{}))
	}
	ok, err := checkFirstLevel(cfg)
	if err != nil {
		return fmt.Errorf("unable to check first level of the config file")
	}
	if ok {
		return nil
	}

	return checkSteps(cfg)
}

// finding executing paths at the first level of the stage
func checkFirstLevel(cfg Config) (bool, error) {
	scr, ok := cfg[script]
	if ok {
		return true, execCommand(scr.(string))
	}

	cmd, ok := cfg[command]
	if ok {
		return true, execCommand(cmd.(string))
	}

	return false, nil
}

// checkSteps provides checking of the steps on config
func checkSteps(cfg Config) error {
	s, ok := cfg[steps]
	if !ok {
		return nil
	}
	for _, i := range s.([]interface{}) {
		st := Stage{}
		if err := mapstructure.Decode(i, &st); err != nil {
			return fmt.Errorf("unable to decode structure: %v", err)
		}
		if err := execStage(st); err != nil {
			return fmt.Errorf("unable to execute stage: %v", err)
		}
	}
	return nil
}

// execStage provides executing of the stage
func execStage(st Stage) error {
	if st.Command == "" {
		return fmt.Errorf("command is not defined")
	}

	if len(st.Envs) > 0 {
		addEnvVariables(st.Envs)
	}

	return execCommand(st.Command)
}

func execCommand(command string) error {
	name, args := prepareCommand(command)
	cmd := exec.Command(name, args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("unable to execute command: %v", err)
	}
	return nil
}

// prepareCommand returns command name and list of args
func prepareCommand(cmd string) (string, []string) {
	res := strings.Split(cmd, " ")
	if len(res) <= 1 {
		return res[0], []string{}
	}
	return res[0], res[1:]
}
