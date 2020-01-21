package parser

import (
	"fmt"
	"github.com/BuldozerAGR/configs"
	"gopkg.in/yaml.v2"
)

func Parse(path string) *configs.BuldozerConfig {
	data, readError := ReadFile(path)
	if readError != nil {
		panic(readError)
	}

	config := configs.BuldozerConfig{}
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- t:\n%v\n\n", config)
	return &config
}
func validate(config configs.BuldozerConfig) {

}
