package main

type Field struct {
	Name string
	Type   string
	Length int
	Minval float64
	Maxval float64
}

type Entity struct {
	Description string
	Fields      []Field
}

type Function struct {
	Description string
	Method string
	Uri string
	Args []Field
}

type Schema struct {
	Entities []Entity
	Functions []Function
}

