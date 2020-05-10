package epit

import (
	"errors"
	"fmt"

	"github.com/saromanov/cowrow"
)

// LoadConfig provides loading of configuration file
func LoadConfig(path string) error {

	if path == "" {
		return errors.New("path to config is empty")
	}

	cfg := map[string]interface{}{}
	err := cowrow.LoadByPath(path, &cfg)
	if err != nil {
		return err
	}

	fmt.Println(cfg)
	return nil
}
