package gync

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func RegisterCommands() []cli.Command {

	return []cli.Command {
		{
			Name: 		"add",
			Aliases: 	nil,
			Usage: 		"Add a new directory/file to be watched",
			Action: func(c *cli.Context) {
				fmt.Println("diretorio adicionado: ", c.Args().First())
			},
		},
	}
}