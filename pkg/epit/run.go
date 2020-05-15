package epit

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// run provides running of the stage
func run(cfg Config) error {
	addEnvVariables(cfg["env"].([]interface{}))
	ok, err := checkFirstLevel(cfg)
	if err != nil {
		return fmt.Errorf("unable to check first level of the config file")
	}
	if ok {
		return nil
	}

	return nil
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
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
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
