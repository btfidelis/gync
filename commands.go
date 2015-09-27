package gync

import (
	//"fmt"
	"github.com/codegangsta/cli"
	"github.com/btfidelis/gync/model"
)

func RegisterCommands() []cli.Command {

	return []cli.Command {
		{
			Name: 		"add",
			Aliases: 	nil,
			Usage: 		"Add a new directory/file to be watched",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				path := c.Args().Get(2)

				save := model.NewSave(name, path)

				if save != nil {
					save.Save()
				}
			},
		},
	}
}