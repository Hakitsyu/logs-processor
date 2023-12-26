package config

import (
	"errors"
	"io/fs"
	"os"
)

type Config struct {
	HelloWorld string `yaml:"helloWorld"`
}

type ConfigInterpreter interface {
	IsValid(fileInfo fs.FileInfo) bool
	Interpret(content []byte) (*Config, error)
}

type ConfigReader struct {
	interpreters *[]ConfigInterpreter
}

func newConfigReader() *ConfigReader {
	interpreters := []ConfigInterpreter{NewYamlInterpreter()}

	return &ConfigReader{
		interpreters: &interpreters,
	}
}

func (reader *ConfigReader) Read(path string) (*Config, error) {
	interpreter, err := reader.getInterpreter(path)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return (*interpreter).Interpret(content)
}

func (reader *ConfigReader) getInterpreter(path string) (*ConfigInterpreter, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	for _, interpreter := range *reader.interpreters {
		if interpreter.IsValid(fileInfo) {
			return &interpreter, nil
		}
	}

	return nil, errors.New("Don't found interpreter for this configuration file")
}

func ReadConfig(path string) (*Config, error) {
	return newConfigReader().Read(path)
}
