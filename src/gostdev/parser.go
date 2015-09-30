package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func parseFieldAttributes(fieldData string, field *Field) error {

	s := strings.Replace(fieldData, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	if len(s) == 0 {
		return errors.New("parse error: empty field data")
	}

	// parse string: `int {0:1} = 0`, `string (255)` etc.
	r, err := regexp.Compile(`^(\w+)(\((\d+)\))?(\{(((\w+)?:(\w+)?)|(\w+([,](\w+))*))\})?(\=(\w+))?$`)
	if err != nil {
		return err
	}

	matches := r.FindStringSubmatch(s)
	if len(matches) == 0 {
		return errors.New("parse error: field regexp not matched")
	}

	field.Type = matches[1]

	if len(matches[3]) > 0 {
		field.Length, err = strconv.Atoi(matches[3])
		if err != nil {
			return err
		}
	}

	if len(matches[7]) > 0 {
		field.Minval, err = strconv.ParseFloat(matches[7], 64)
		if err != nil {
			return err
		}
	}

	if len(matches[8]) > 0 {
		field.Maxval, err = strconv.ParseFloat(matches[8], 64)
		if err != nil {
			return err
		}
	}

	return nil
}
