package app 

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/btfidelis/gync/model"
)

var STARTDAEMON bool
var DROPBOX_LOCAL string


func Boot() {

	DROPBOX_LOCAL = model.GetConfig().BackupPath

	app := cli.NewApp()
	app.Name = "Gync"
	app.Usage = "Keep files and directories synced with in real time"
	app.Commands = RegisterCommands();
	app.Run(os.Args)
}



