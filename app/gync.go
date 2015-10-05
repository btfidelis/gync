package app 

import (
	"os"
	"github.com/codegangsta/cli"
)

func Boot() {
	app := cli.NewApp()
	app.Name = "Gync"
	app.Usage = "Keep files and directories synced with in real time"

	app.Commands = RegisterCommands();

	app.Run(os.Args)
}

