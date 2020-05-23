package main

import (
	"fmt"
	"log"
	"os"

	"github.com/saromanov/epit/pkg/epit"
	"go.uber.org/zap"
)

const configFilePath = ".epit.yml"

func main() {
	stage := os.Getenv("STAGE")
	args := os.Args
	if stage == "" && len(args) == 1 {
		log.Fatal("name of the stage is not defined")
	}
	if stage == "" {
		stage = args[1]
	}
	log, err := zap.NewProduction()
	if err != nil {
		panic("unable to init logging")
	}
	if err := epit.ExecStage(log, configFilePath, stage); err != nil {
		log.Fatal(fmt.Sprintf("unable to parse config: %v", err))
	}
}
