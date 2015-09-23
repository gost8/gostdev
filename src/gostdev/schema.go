package main

type FieldAttributes struct {
	Type   string
	Length int
	Minval float64
	Maxval float64
}

type Entity struct {
	Description string
	Fields      []map[string]string
}

type Schema struct {
	Entities map[string]Entity
}
