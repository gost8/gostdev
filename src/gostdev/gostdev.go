package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Field struct {
	Name string
	Type string
}

type Fields []Field

type Entity struct {
	Fields Fields
}

type Schema struct{
	Entity Entity
}

func main() {
	app := cli.NewApp()
	app.Name = "gostdev"
	app.Usage = "Golang code generation tool"
	app.Author = "Yuri Karamani <y.karamani@gmail.com>"
	app.Commands = []cli.Command{
		{
			Name: "create",
			Subcommands: []cli.Command{
      			{
 					Name:  "routes",
 					Usage: "create routes for web application",
					Action: func(c *cli.Context) {
						data, err := ioutil.ReadFile("/home/y2k/projects/gostdev/test/fixtures/schema1/schema.yaml")
						if err != nil {
							fmt.Printf("error: %v", err)
							panic(err)
						}
						
						fmt.Print(string(data))
						entity := Schema{}
						err = yaml.Unmarshal(data, &entity)
						if err != nil {
							fmt.Printf("error: %v", err)
							panic(err)
						}
						fmt.Printf("--- t:\n%v\n\n", entity)
						fmt.Println("create routes")
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