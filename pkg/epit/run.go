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
func run(name string, cfg Config) error {
	envs, okEvs := cfg[env]
	if okEvs {
		prepareEnvVars(envs.([]interface{}), setEnvVariables)
	}
	ok, err := checkFirstLevel(name, cfg)
	if err != nil {
		return fmt.Errorf("unable to check first level of the config file")
	}
	if ok {
		if okEvs {
			prepareEnvVars(envs.([]interface{}), unsetEnvVariables)
		}
		return nil
	}

	return checkSteps(cfg)
}

// finding executing paths at the first level of the stage
func checkFirstLevel(name string, cfg Config) (bool, error) {
	info("Executing of the task %s\n", name)
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
	for i, n := range s.([]interface{}) {
		st := Stage{}
		if err := mapstructure.Decode(n, &st); err != nil {
			return fmt.Errorf("unable to decode structure: %v", err)
		}
		step := st.Name
		if step == "" {
			step = fmt.Sprintf("%d", i+1)
		}
		info("Executing of the step %s\n", step)
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
		prepareEnvVars(st.Envs, setEnvVariables)
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
