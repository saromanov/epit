package main

import (
	"fmt"
	"log"
	"os"

	"github.com/saromanov/epit/pkg/epit"
	"go.uber.org/zap"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("name of the stage is not defined")
	}
	log, err := zap.NewProduction()
	if err != nil {
		panic("unable to init logging")
	}
	if err := epit.ExecStage("./examples/basic/.epit.yml", args[1]); err != nil {
		log.Fatal(fmt.Sprintf("unable to parse config: %v", err))
	}
}
