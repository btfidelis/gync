package gync

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/btfidelis/gync/model"
	"github.com/crackcomm/clitable"
)

func RegisterCommands() []cli.Command {

	return []cli.Command {
		{
			Name: 		"add",
			Aliases: 	nil,
			Usage: 		"Add a new directory/file to be watched",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				path := c.Args().Get(1)

				save := model.NewSave(name, path)

				if save != nil {
					save.Save()
				} else {
					fmt.Println("Unable to add game")
				}
			},
		},
		{
			Name:	    "list",
			Aliases:    []string{"ls"},
			Usage:		"Lists regitered saves",
			Action:	func (c *cli.Context) {
				col := model.GetSaveCollection()
				table := clitable.New([]string{"Name", "Path"});

				for _, save := range(col.Saves) {
					table.AddRow(map[string]interface{}{"Name": save.Name, "Path": save.Location })
				}

				table.Print()
			},
		},
	}
}