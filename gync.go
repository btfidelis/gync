package main 

import(
	"time"
	"fmt"
	"os"
	"github.com/btfidelis/gync/app"
	"github.com/btfidelis/gync/model"
	"github.com/btfidelis/gync/core"
)

func main() {
	app.Boot()

	if !app.STARTDAEMON {
		os.Exit(0)
	}

	core.InitNotifyManager()

	saveChange := make(chan bool)
	saveCol := model.GetSaveCollection()

	saveFileWatcher := app.Watcher{
		ModTimes: make(map[string]time.Time, 0),
		ModFiles: make(map[string]int, 0),
	}

	go saveFileWatcher.ObserveFile(model.GetSaveLocal(), saveChange)

	fileWatchers := make([]app.Watcher, len(saveCol.Saves))
	
	for i, _:= range(fileWatchers) {
		fileWatchers[i] = app.Watcher{
			ModTimes: make(map[string]time.Time, 0),
			ModFiles: make(map[string]int, 0),
			Dir:	  saveCol.Saves[i].Location,
			Root:	  saveCol.Saves[i].Name,
		}

		go fileWatchers[i].ObserveDir()
	}
	

	for {
		if <-saveChange {
			fmt.Println("Please restart the deamon if you want to commit the changes")
		}
	}
}