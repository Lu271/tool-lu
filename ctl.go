package main

import (
	"github.com/Lu271/tool-lu/crud"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var commands = []*cli.Command{
	{
		Name:  "crud",
		Usage: "generate DB CRUD code",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "t",
				Usage: "db type",
			},
			&cli.StringFlag{
				Name:  "n",
				Usage: "db dsn",
			},
			&cli.StringFlag{
				Name:  "f",
				Usage: "folder",
			},
		},
		Action: crud.Init(),
	},
}

func main() {
	app := cli.NewApp()
	app.Usage = "a cli tool to generate code"
	app.Commands = commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
