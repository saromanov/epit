package epit

import (
	"os"
	"strings"
)

// adding environment variables to the stage
func addEnvVariables(vars []interface{}) {
	if len(vars) == 0 {
		return
	}

	for _, v := range vars {
		s := v.(string)
		data := strings.Split(strings.TrimSpace(s), "=")
		if len(data) != 2 {
			continue
		}
		os.Setenv(data[0], data[1])
	}
}
