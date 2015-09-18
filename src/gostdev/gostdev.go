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
	app.Action = func(c *cli.Context) {
		fmt.Println("Cool tool!")
	}

	app.Run(os.Args)
}