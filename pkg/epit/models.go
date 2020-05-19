package epit

type Stage struct {
	// Command for execution
	Command string
	// Name of the step
	Name string
	// Show duration of execution
	Duration bool
	// list of environment variables in the format FOO=BAR
	Envs []interface{}
}

// Config provides definition of configuration
type Config map[string]interface{}
