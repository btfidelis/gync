package gync 

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	"github.com/stacktic/dropbox"
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

func main() {

	file, err := os.Open("storage/dropbox.json")

	if err != nil {
		log.Fatal(err)
	}

	jsonInfo := json.NewDecoder(file)

	if err = jsonInfo.Decode(&DropboxInfo); err != nil {
		log.Fatal(err)
	}

	dropBox := dropbox.NewDropbox()
	dropBox.SetAppInfo(DropboxInfo.ClientId, DropboxInfo.ClientSecret)
	dropBox.SetAccessToken(DropboxInfo.Token)

	app := cli.NewApp()
	app.Name = "Gync"
	app.Usage = "Keep files and directories synced with in real time"

	app.Commands = RegisterCommands();

	app.Run(os.Args)


  /*if _, err := dropBox.CreateFolder("UnoMas"); err != nil {
		fmt.Printf("Error creating folder %s \n", err)
	} else {
		fmt.Println("Folder created !")
	}*/
}
