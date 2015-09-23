package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func parseFieldAttributes(fieldData string) (*FieldAttributes, error) {

	attrs := &FieldAttributes{}

	s := strings.Replace(fieldData, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	if len(s) == 0 {
		return nil, errors.New("parse error: empty field data")
	}

	// parse string: `int {0:1} = 0`, `string (255)` etc.
	r, err := regexp.Compile(`^(\w+)(\((\d+)\))?(\{(((\w+)?:(\w+)?)|(\w+([,](\w+))*))\})?(\=(\w+))?$`)
	if err != nil {
		return nil, err
	}

    matches := r.FindStringSubmatch(s)
	if len(matches) == 0 {
		return nil, errors.New("parse error: field regexp not matched")
	}

	attrs.Type = matches[1]

	if len(matches[3]) > 0 {
		attrs.Length, err = strconv.Atoi(matches[3])
		if err != nil {
			return nil, err
		}
	}

	if len(matches[7]) > 0 {
		attrs.Minval, err = strconv.ParseFloat(matches[7], 64)
		if err != nil {
			return nil, err
		}
	}

	if len(matches[8]) > 0 {
		attrs.Maxval, err = strconv.ParseFloat(matches[8], 64)
		if err != nil {
			return nil, err
		}
	}

	return attrs, nil
}
