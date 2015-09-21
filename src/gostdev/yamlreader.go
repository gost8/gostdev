package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func loadYamlSchema(schemaFile string) ([]byte, error) {

	file := "schema.yaml"
	if len(schemaFile) > 0 {
		file = schemaFile
	}

	return ioutil.ReadFile(file)
}

func unmarshalYamlSchema(data []byte) (*Schema, error) {
	schema := Schema{}
	err := yaml.Unmarshal(data, &schema)
	return &schema, err
}
