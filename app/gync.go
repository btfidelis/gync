package app 

import (
	"os"
	"github.com/codegangsta/cli"
	"flag"
)

var STARTDAEMON bool

func Boot() {
	flag.BoolVar(&STARTDAEMON, "daemon", false, "start gync deamon")

	app := cli.NewApp()
	app.Name = "Gync"
	app.Usage = "Keep files and directories synced with in real time"
	app.Commands = RegisterCommands();
	app.Run(os.Args)

	flag.Parse()
}

