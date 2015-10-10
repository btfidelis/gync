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
	Dir 		string
	Root		string
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
		fmt.Println(w.ModFiles, "path: ", path, " modified")	
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

func (w *Watcher) ObserveDir(save model.Save) {
	w.ModTimes = make(map[string]time.Time, 0)
	
	error := filepath.Walk(save.Location, w.WalkPopulate)
	mirror := new(Watcher)
	
	
	if error != nil {
		log.Println("ON DAEMON: ", error)
		return
	}

	mirror.ModTimes = make(map[string]time.Time, 0)
	error = filepath.Walk(save.Location, mirror.WalkPopulate)

	if error != nil {
		log.Println("ON DAEMON: ", error)
		return	
	}
	
	for {
		err := filepath.Walk(save.Location, w.WalkCheck)
		
		if err != nil {
			log.Println("ON DAEMON: ", err)
			return
		}

		time.Sleep(time.Second * 2)
	}
}

func (w *Watcher) sync(path string) {
	// if file was renamed, if file was moved, if file was deleted
	backup := model.GetConfig().BackupPath
	fmt.Println("syncing: ", path)
	i, err := os.Stat(path);

	if err != nil {
		log.Println("Error Sync: ", err)
	}


	if i.IsDir() {

		destination, err := filepath.Rel(w.Dir, path)

		if err != nil {
			log.Println(err)
		}

		if _, err = os.Stat(filepath.Join(filepath.Join(backup, w.Root), destination)); os.IsNotExist(err) {
			
			err := os.Mkdir(filepath.Join(filepath.Join(backup, w.Root), destination), 0777)

			if err != nil {
				log.Println("New dir: ",  err)
			}
		} 

		for dir, _:= range(w.ModTimes) {

			if _,err := os.Stat(dir); os.IsNotExist(err) {
				fmt.Println("moved or renamed: ", dir)
				delete(w.ModTimes, dir)
				w.ModFiles[dir] = DELETED
			}
		}
	} else {
		fmt.Println(w.ModFiles, "path: ", path)
		w.ModFiles[path] = MODIFIED
	}

	go w.copy(backup)

}

func (w *Watcher) copy(dest string) {
	for path, val := range(w.ModFiles) {

		destination, err := filepath.Rel(w.Dir, path)
		
		switch val {
			case MODIFIED:

				if err != nil {
					log.Println("error sync: ", err)
				}

				file, err := os.Stat(path)

				if err != nil {
					log.Println("error sync: ", err)
				}

				if file.IsDir() {
					err := os.Mkdir(filepath.Join(filepath.Join(dest, w.Root), destination), 0777)
					fmt.Println("coping dir: ", filepath.Join(filepath.Join(dest, w.Root), destination))
				
					if err != nil {
						log.Println("error sync: ", err)
					}
				} else {
					err := core.CopyFile(path, filepath.Join(filepath.Join(dest, w.Root), destination))
					
					if err != nil {
						log.Println("error sync: ", err)
					}
				}

				break

			case DELETED:
				err := os.Remove(filepath.Join(filepath.Join(dest, w.Root), destination))
				
				if err != nil {
					log.Println("error sync: ", err)
				}
				break

			default:
		}

		delete(w.ModFiles, path)
	}
}
