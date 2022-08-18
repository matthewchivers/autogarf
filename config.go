package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type conf struct {
	ClientDir string `yaml:"client-directory"`
}

// Reads the config file and returns the config struct
func readConfig(filename string) (*conf, error) {
	log.Printf("Reading config file: %s", filename)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &conf{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	log.Printf("Finished reading config file")
	return config, nil
}
