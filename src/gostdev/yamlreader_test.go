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

	entity, ok := schema.Entities["User"]
	if !ok {
		t.Errorf("Entity not found! unmarshalYamlSchema(%q) == %q", YamlSchema1, schema)
	}

	if len(entity.Fields) == 0 {
		t.Errorf("Fields is empty! unmarshalYamlSchema(%q) == %q", YamlSchema1, schema)
	}

	fieldId, ok := entity.Fields[0]["id"]
	if !ok {
		t.Errorf("Field not found! unmarshalYamlSchema(%q) == %q", YamlSchema1, schema)
	}
	if fieldId != "int" {
		t.Errorf("Field data error! unmarshalYamlSchema(%q) == %q", YamlSchema1, schema)
	}

}
