package main

import (
	"github.com/BuldozerAGR/cmd/generator"
	"github.com/BuldozerAGR/cmd/parser"
)

func main() {
	config := parser.Parse("buldozer.yaml")
	generator.StepAlgorithm(&config.SiteSetup.Schedule.StepLoad)
}
