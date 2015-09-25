package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YEntity struct {
	Description string
	Fields      map[string]string
}

type YFunction struct {
	Description string
	Method string
	Uri string
	Args map[string]string
}

type YSchema struct {
	Entities map[string]YEntity
	Functions map[string]YFunction
}

func loadYamlSchema(schemaFile string) ([]byte, error) {

	file := "schema.yaml"
	if len(schemaFile) > 0 {
		file = schemaFile
	}

	return ioutil.ReadFile(file)
}

func unmarshalYamlSchema(data []byte) (*YSchema, error) {
	schema := YSchema{}
	err := yaml.Unmarshal(data, &schema)
	return &schema, err
}
