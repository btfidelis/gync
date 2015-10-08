package app

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
			Usage: 		"Adds a new directory/file to be watched",
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
			Usage:		"Lists regitered directories",
			Action:	func (c *cli.Context) {
				col := model.GetSaveCollection()
				table := clitable.New([]string{"Name", "Path"})

				for _, save := range(col.Saves) {
					table.AddRow(map[string]interface{}{"Name": save.Name, "Path": save.Location })
				}

				table.Print()
			},
		},
		{
			Name:	    "remove",
			Aliases:    []string{"rm"},
			Usage:		"Removes a registered directory",
			Action:	func (c *cli.Context) {
				col := model.GetSaveCollection()
				name := c.Args().First()

				save, id := col.Where(name)

				if save != nil {
					col.Remove(id)
				} else {
					fmt.Println("Failed deleting, unable to find the game: ", name)
				}
			},
		},
		{
			Name:	    "start",
			Aliases:    nil,
			Usage:		"Starts gync as a daemon",
			Action:	func (c *cli.Context) {
				fmt.Println("Starting daemon...")
				STARTDAEMON = true
			},
		},
	}
}
