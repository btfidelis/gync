package app

import (
	"os"
	"fmt"
	"log"
	"time"
	"path/filepath"
	"github.com/btfidelis/gync/model"
	"github.com/btfidelis/gync/core"
)

const (
	DELETED 	=	iota
	MODIFIED	=	iota
)

type Watcher struct {
	ModTimes 	map[string]time.Time
	ModFiles	map[string]int
}

/**
 * Observe file
 */
func (w Watcher) ObserveFile(file string, changed chan bool) {
	info, err := os.Stat(file)

	if err == nil {
		lastChangedOn := info.ModTime()

		for {
		
			info, err := os.Stat(file)
			
			if err == nil {

				if info.ModTime().After(lastChangedOn) {
					lastChangedOn = info.ModTime()
					changed <- true
				}
			}

			changed <- false
			time.Sleep(time.Second * 2)
		}
	}
}

func (w *Watcher) WalkCheck(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Println(err)
		return err
	}

	if ! w.ModTimes[path].Equal(info.ModTime()) {
		fmt.Println(info.Name(), "coping files :", path)
		w.ModTimes[path] = info.ModTime()
		
		w.sync(path)
	}

	return err
}

func (w *Watcher) WalkPopulate(path string, info os.FileInfo, err error) error {

	if err != nil {
		log.Println(err)
		return err
	}
	
	w.ModTimes[path] = info.ModTime()

	return err
}

func (w *Watcher) ObserveDir(path string) {
	w.ModTimes = make(map[string]time.Time, 0)
	
	error := filepath.Walk(path, w.WalkPopulate)
	mirror := new(Watcher)
	
	
	if error != nil {
		log.Println("ON DAEMON: ", error)
		return
	}

	mirror.ModTimes = make(map[string]time.Time, 0)
	error = filepath.Walk(path, mirror.WalkPopulate)

	if error != nil {
		log.Println("ON DAEMON: ", error)
		return	
	}
	
	for {
		err := filepath.Walk(path, w.WalkCheck)
		
		if err != nil {
			log.Println("ON DAEMON: ", err)
			return
		}

		time.Sleep(time.Second * 2)
	}
}

func (w *Watcher) sync(path string) {
	// if file was renamed, if file was moved, if file was deleted
	fmt.Println("syncing: ", path)
	i, err := os.Stat(path);

	if err != nil {
		log.Println("Error Sync: ", err)
	}

	if i.IsDir() {	
		for dir, _:= range(w.ModTimes) {

			if _,err := os.Stat(dir); os.IsNotExist(err) {
				fmt.Println("moved or renamed: ", dir)
				delete(w.ModTimes, dir)
				w.ModFiles[dir] = DELETED
			}
		}
	} else {
		w.ModFiles[path] = MODIFIED
	}

	go w.copy()

	fmt.Println(w.ModFiles)
}

func (w *Watcher) copy() {
	for path, val := range(w.ModFiles) {

		switch val {

			case MODIFIED:
				err := core.CopyFile(path, filepath.Join(model.GetConfig().BackupPath, filepath.Clean(path)))
				if err != nil {
					log.Println("error sync: ", err)
				}
				break
			case DELETED:
				err := os.Remove(path)
				if err != nil {
					log.Println("error sync: ", err)
				}
				break
			default:
		}

		delete(w.ModFiles, path)
	}
}
