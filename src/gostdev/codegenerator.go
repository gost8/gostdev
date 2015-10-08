package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type CodeGenerator struct {
	TemplateFile string
	Schema *Schema
}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

func (c *CodeGenerator) SetTemplateFile(templateFile string) *CodeGenerator {
	c.TemplateFile = templateFile
	return c
}

func (c *CodeGenerator) SetSchema(schema *Schema) *CodeGenerator {
	c.Schema = schema
	return c
}

func (c *CodeGenerator) Generate() {

	funcMap := template.FuncMap{
		"split": splitDescription,
	}

	templateData, _ := ioutil.ReadFile(c.TemplateFile)
	t, err := template.New("code").Funcs(funcMap).Parse(string(templateData))
	if err != nil {
		panic(err)
	}
	
	err = t.Execute(os.Stdout, c.Schema.Functions)
	if err != nil {
		panic(err)
	}
}
