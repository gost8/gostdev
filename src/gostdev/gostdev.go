package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

type GlobalFlags struct {
	Schema  string
	Verbose bool
}

var globalFlags GlobalFlags

func parseGlobalFlags(c *cli.Context) {
	globalFlags.Schema = c.String("schema")
	globalFlags.Verbose = c.Bool("verbose")
}

func main() {
	app := cli.NewApp()
	app.Name = "gostdev"
	app.Usage = "Golang code generation tool"
	app.Author = "Yuri Karamani <y.karamani@gmail.com>"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "schema, s",
			Usage: "schema for generation",
		},
	}

	app.Before = func(c *cli.Context) error {
		parseGlobalFlags(c)
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name: "create",
			Subcommands: []cli.Command{
				{
					Name:  "routes",
					Usage: "create routes for web application",
					Action: func(c *cli.Context) {
						fmt.Println("create routes")

						data, err := loadYamlSchema(globalFlags.Schema)
						if err != nil {
							fmt.Printf("error: %v", err)
							panic(err)
						}
						fmt.Printf("data:\n%v\n\n", string(data))

						yschema, err := unmarshalYamlSchema(data)
						if err != nil {
							fmt.Printf("Yaml error: %v", err)
							panic(err)
						}

						schema := &Schema{}

						for entityName, entityData := range yschema.Entities {
							entity := NewEntity().
								SetDescription(entityData.Description).
								SetName(entityName)
							fmt.Printf("%s:\n", entityName)
							for fieldName, fieldData := range entityData.Fields {
								fmt.Printf("%s: %s\n", fieldName, fieldData)
								field := &Field{Name: fieldName}
								err := parseFieldAttributes(fieldData, field)
								if err != nil {
									fmt.Printf("Parse error: %v", err)
									panic(err)
								}
								fmt.Printf("%v\n", field)
								entity.AddField(field)
							}
							schema.AddEntity(entity)
						}

						for functionName, functionData := range yschema.Functions {
							function := NewFunction().
								SetDescription(functionData.Description).
								SetName(functionName).
								SetMethod(functionData.Method).
								SetUri(functionData.Uri)
							fmt.Printf("%s:\n", functionName)
							for argName, argData := range functionData.Args {
								fmt.Printf("%s: %s\n", argName, argData)
								argField := &Field{Name: argName}
								err := parseFieldAttributes(argData, argField)
								if err != nil {
									fmt.Printf("Parse error: %v", err)
									panic(err)
								}
								fmt.Printf("%v\n", argField)
								function.AddArg(argField)
							}
							schema.AddFunction(function)
						}

						generateRoutes(schema)
						generateOperations(schema)
					},
				},
				{
					Name:  "models",
					Usage: "create models for web application",
					Action: func(c *cli.Context) {
						fmt.Println("create models")
					},
				},
				{
					Name:  "webclient",
					Usage: "create client package for web application",
					Action: func(c *cli.Context) {
						fmt.Println("create webclient")
					},
				},
			},
		},
	}

	app.Action = func(c *cli.Context) {
		fmt.Println("Cool tool!")
	}

	app.Run(os.Args)
}
