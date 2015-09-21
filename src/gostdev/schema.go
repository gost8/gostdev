package main

type Entity struct {
	Fields []map[string]string
}

type Schema struct {
	Entities map[string]Entity
}
