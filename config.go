package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type conf struct {
	ClientDir string `yaml:"client-directory"`
}

// Reads the config file and returns the config struct
func readConfig(filename string) (*conf, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &conf{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
