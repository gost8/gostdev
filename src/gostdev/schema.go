package main

type FieldAttributes struct {
	Type string
	Length int
}

type Entity struct {
	Fields []map[string]string
}

type Schema struct {
	Entities map[string]Entity
}
