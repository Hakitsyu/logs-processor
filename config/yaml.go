package config

import (
	"io/fs"
	"strings"

	"gopkg.in/yaml.v3"
)

type YamlInterpreter struct {
}

func (interpreter *YamlInterpreter) IsValid(fileInfo fs.FileInfo) bool {
	return strings.HasSuffix(fileInfo.Name(), ".yaml")
}

func (interpreter *YamlInterpreter) Interpret(content []byte) (*Config, error) {
	config := Config{}
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func NewYamlInterpreter() *YamlInterpreter {
	return &YamlInterpreter{}
}
