package main

import (
	"errors"
	"strings"
)

func parseFieldAttributes(fieldData string) (*FieldAttributes, error) {

	attrs := &FieldAttributes{}

	s := strings.Trim(fieldData, " ")
	if len(s) == 0 {
		return nil, errors.New("parse error: empty field data")
	}

	idxTypeEnd := strings.IndexAny(s, " ({=\n")
	if idxTypeEnd == 0 {
		return nil, errors.New("parse error: unknown type")
	}

	if idxTypeEnd == -1 {
		attrs.Type = s
	} else {
		attrs.Type = s[0:idxTypeEnd]
	}
	if len(attrs.Type) == 0 {
		return nil, errors.New("parse error")
	}

	return attrs, nil
}
