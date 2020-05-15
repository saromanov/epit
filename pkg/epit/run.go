package epit

import (
	"fmt"
)

// run provides running of the stage
func run(cfg Config) error {

}

// finding executing paths at the first level of the stage
func checkFirstLevel(cfg Config) (bool, error) {
	script, ok := cfg["script"]
	if ok {
		fmt.Println("SCRIPT: ", script)
		return
	}
}
