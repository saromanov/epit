package epit

type Stage struct {
	Command string
	Name    string
	Envs    []interface{}
}

// Config provides definition of configuration
type Config map[string]interface{}
