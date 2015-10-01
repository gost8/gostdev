package main

import "strings"

func splitDescription(descr string) []string {
	s := strings.Trim(descr, "\n")
	return strings.Split(s, "\n")
}
