package app 

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/btfidelis/gync/model"
)

var STARTDAEMON bool

// Global Options
var DROPBOX_LOCAL string
var CHECK_INTERVAL int


func Boot() {
	conf := model.GetConfig()
	
	DROPBOX_LOCAL = conf.BackupPath
	CHECK_INTERVAL = conf.CheckInterval

	model.CheckSaveFile()

	app := cli.NewApp()
	app.Name = "Gync"
	app.Usage = "Keep files and directories synced in real time"
	app.Commands = RegisterCommands();
	app.Run(os.Args)
}