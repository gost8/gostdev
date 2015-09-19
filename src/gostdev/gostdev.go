package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

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