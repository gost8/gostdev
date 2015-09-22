package main

import (
	"testing"
)

func TestParseFieldAttributesSuccess(t *testing.T) {
	cases := []struct {
		in string
		want FieldAttributes
	}{
		{"int", FieldAttributes{Type: "int"}},
		{" int ", FieldAttributes{Type: "int"}},
		{" int\n", FieldAttributes{Type: "int"}},
		{"string (255)", FieldAttributes{Type: "string"}},
		{"string(255)", FieldAttributes{Type: "string"}},
		{"int{1:255}", FieldAttributes{Type: "int"}},
		{"int\n", FieldAttributes{Type: "int"}},
		{"int=20", FieldAttributes{Type: "int"}},
	}

	for _, c := range cases {
		attrs, err := parseFieldAttributes(c.in)
		if err != nil {
			t.Error(c.in, err)
		}
		if *attrs != c.want {
			t.Errorf("parseFieldAttributes(%q) == %q, want %q", c.in, attrs, c.want)
		}
	}
}

func TestParseFieldAttributesError(t *testing.T) {
	cases := []struct {
		in string
		wantErrorText string
	}{
		{"", "parse error: empty field data"},
		{" ", "parse error: empty field data"},
		{"()", "parse error: unknown type"},
		{" {()} int", "parse error: unknown type"},
	}

	for _, c := range cases {
		attrs, err := parseFieldAttributes(c.in)
		if attrs != nil {
			t.Errorf("parseFieldAttributes(%q) == %q, want nil", c.in, attrs)
		}
		if err == nil {
			t.Errorf("parseFieldAttributes(%q) return not error, want '%q'", c.in, c.wantErrorText)
		} else {
			if err.Error() != c.wantErrorText {
				t.Errorf("parseFieldAttributes(%q) return error '%q', want '%q'", c.in, err.Error(), c.wantErrorText)
			}
		}
	}
}
