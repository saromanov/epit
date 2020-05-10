package epit

import "fmt"

// ExecStage provides execution of the stage
func ExecStage(path, name string) error {
	cfg, err := loadConfig(path)
	if err != nil {
		return fmt.Errorf("ExecStage: unable to load config: %v", err)
	}
	stage, ok := cfg[name]
	if !ok {
		return fmt.Errorf("name of the stahe is not found")
	}
	fmt.Println(stage)
	return nil
}