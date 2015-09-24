package main

import (
	"testing"
)

const YamlSchema1 = `
entities:
  User:
    description: |
      This is example
      multiline text
      in description
    fields:
      - id: int

functions:
  getUser:
    description: Get one user by id
    method: GET
    uri: /users/{userId}
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

	if entity.Description != "This is example\nmultiline text\nin description\n" {
		t.Errorf("Description invalid!\n%#v", entity.Description)
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

	if len(schema.Functions) != 1 {
		t.Errorf("unmarshalYamlSchema(%q) == %q", YamlSchema1, schema)
	}

	function, ok := schema.Functions["getUser"]
	if !ok {
		t.Errorf("Function not found! unmarshalYamlSchema(%q) == %q", YamlSchema1, schema)
	}

	if function.Description != "Get one user by id" ||
		function.Method != "GET" ||
		function.Uri != "/users/{userId}" {

		t.Errorf("Function штмфдшв! %q", function)
	}
}
