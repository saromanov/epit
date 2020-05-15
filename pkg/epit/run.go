package epit

import (
	"fmt"
)

// run provides running of the stage
func run(cfg Config) error {
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
		fmt.Println("SCRIPT: ", scr)
		return true, nil
	}

	cmd, ok := cfg[command]
	if ok {
		fmt.Println("cmd ", cmd)
		return true, nil
	}

	return false, nil
}
