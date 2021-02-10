package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/saromanov/epit/pkg/epit"
	"go.uber.org/zap"
)

func main() {
	configFilePath := flag.String("config", ".epit.yml", "path to config")
	flag.Parse()
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
		log.Fatal(fmt.Sprintf("unable to init logging: %v", err))
	}
	if err := epit.Exec(log, *configFilePath, stage); err != nil {
		log.Fatal(fmt.Sprintf("unable to parse config: %v", err))
	}
}
