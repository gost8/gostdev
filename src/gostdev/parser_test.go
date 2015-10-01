package main

import (
	"testing"
)

func TestParseFieldAttributesSuccess(t *testing.T) {
	cases := []struct {
		in   string
		want Field
	}{
		{"string ( 255 ) { 1 : 2 } = 11", Field{Type: "string", Length: 255, Minval: 1, Maxval: 2}},
		{"int", Field{Type: "int"}},
		{" int ", Field{Type: "int"}},
		{" int\n", Field{Type: "int"}},
		{"string (25)", Field{Type: "string", Length: 25}},
		{"string\t(25)", Field{Type: "string", Length: 25}},
		{"int{1:255}", Field{Type: "int", Minval: 1, Maxval: 255}},
		{"int=20", Field{Type: "int"}},
	}

	for _, c := range cases {
		field := &Field{}
		err := parseFieldAttributes(c.in, field)
		if err != nil {
			t.Error(c.in, err)
		}
		if *field != c.want {
			t.Errorf("parseFieldAttributes(%v) == %v, want %v", c.in, field, c.want)
		}
	}
}

func TestParseFieldAttributesError(t *testing.T) {
	cases := []struct {
		in            string
		wantErrorText string
	}{
		{"", "parse error: empty field data"},
		{" ", "parse error: empty field data"},
		{"()", "parse error: field regexp not matched"},
		{" {()} int", "parse error: field regexp not matched"},
	}

	for _, c := range cases {
		field := &Field{}
		err := parseFieldAttributes(c.in, field)
		if err == nil {
			t.Errorf("parseFieldAttributes(%v) return not error, want '%v'", c.in, c.wantErrorText)
		} else {
			if err.Error() != c.wantErrorText {
				t.Errorf("parseFieldAttributes(%q) return error '%v', want '%v'", c.in, err.Error(), c.wantErrorText)
			}
		}
	}
}
