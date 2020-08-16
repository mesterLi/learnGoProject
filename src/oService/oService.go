package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name: "oService",
		Usage: "input 'oService login' to login",
		Action: func(context *cli.Context) error {
			fmt.Print("context")
			return nil
		},
	}
	app.Run(os.Args)
}
