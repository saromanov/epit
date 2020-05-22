package epit

import (
	"os"
	"strings"
)

func prepareEnvVars(vars []interface{}, action func(string, string)) {
	if len(vars) == 0 {
		return
	}

	for _, v := range vars {
		if v == nil {
			continue
		}
		s := v.(string)
		if s == "" {
			continue
		}
		data := strings.Split(strings.TrimSpace(s), "=")
		if len(data) != 2 {
			continue
		}
		action(data[0], data[1])
	}
}

// setting environment variables to the stage
func setEnvVariables(k, v string) {
	os.Setenv(k, v)
}

func unsetEnvVariables(k, v string) {
	os.Unsetenv(k)
}
