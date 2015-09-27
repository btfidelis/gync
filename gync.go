package gync 

import (
	"os"
	"fmt"
	"github.com/codegangsta/cli"
)

var DropboxInfo struct {
	ClientId string
	ClientSecret string
	Token string
}


func listenOn(file string) {
	info, err := os.Stat(file)

	if err == nil {
		lastChangedOn := info.ModTime()

		for {
		
			info, err := os.Stat(file)
			
			if err == nil {

				if info.ModTime().After(lastChangedOn) {
					fmt.Println("arquivo alterado")
					lastChangedOn = info.ModTime()
				}
			}
		}
	}

}

func Boot() {
	app := cli.NewApp()
	app.Name = "Gync"
	app.Usage = "Keep files and directories synced with in real time"

	app.Commands = RegisterCommands();

	app.Run(os.Args)


}

