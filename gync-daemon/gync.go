package main 

import(
	"time"
	"fmt"
	"os"
	"github.com/btfidelis/gync/app"
	"github.com/btfidelis/gync/model"
)

func main() {
	app.Boot()

	if !app.STARTDAEMON {
		os.Exit(0)
	}

	saveChange := make(chan bool)
	saveCol := model.GetSaveCollection()

	saveFileWatcher := app.Watcher{
		ModTimes: make(map[string]time.Time, 0),
		ModFiles: make(map[string]int, 0),
	}

	go saveFileWatcher.ObserveFile(model.GetSaveLocal(), saveChange)

	fileWatchers := make([]app.Watcher, len(saveCol.Saves))
	
	for i, watch := range(fileWatchers) {
		go watch.ObserveDir(saveCol.Saves[i].Location)
	}
	
	for {
		if <-saveChange {
			fmt.Println("changed")
		}
	}
}