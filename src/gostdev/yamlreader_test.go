package main

import (
	"testing"
)

const YamlSchema1 = `
entities:
  User:
    fields:
      - id: int
`

func TestUnmarshalYamlSchema1(t *testing.T) {

	schema, err := unmarshalYamlSchema([]byte(YamlSchema1))
	if err != nil {
		t.Error(YamlSchema1, err)
	}
	if len(schema.Entities) != 1 {
		t.Errorf("unmarshalYamlSchema(%q) == %q", YamlSchema1, schema)
	}
}
