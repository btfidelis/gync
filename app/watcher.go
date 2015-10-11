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
	NEWDIR		=   iota
)

type Watcher struct {
	ModTimes 	map[string]time.Time
	ModFiles	map[string]int
	Dir 		string
	Root		string
}

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

func (w *Watcher) walkCheck(path string, info os.FileInfo, err error) error {
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

func (w *Watcher) walkPopulate(path string, info os.FileInfo, err error) error {

	if err != nil {
		log.Println(err)
		return err
	}
	
	w.ModTimes[path] = info.ModTime()

	return err
}

func (w *Watcher) ObserveDir(save model.Save) {
	w.ModTimes = make(map[string]time.Time, 0)
	
	error := filepath.Walk(save.Location, w.walkPopulate)
	mirror := new(Watcher)
	
	
	if error != nil {
		log.Println("ON DAEMON: ", error)
		return
	}

	mirror.ModTimes = make(map[string]time.Time, 0)
	error = filepath.Walk(save.Location, mirror.walkPopulate)

	if error != nil {
		log.Println("ON DAEMON: ", error)
		return	
	}
	
	for {
		err := filepath.Walk(save.Location, w.walkCheck)
		
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
		
			w.ModFiles[path] = NEWDIR
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

		relBackupPath, err := filepath.Rel(w.Dir, path)

		if err != nil {
			log.Println("error sync: ", err)
		}
		
		destPath := filepath.Join(filepath.Join(dest, w.Root), relBackupPath)
		
		switch val {
			case NEWDIR:
				err := os.Mkdir(destPath, 0777)

				if err != nil {
					log.Println("New dir: ",  err)
				}
			break

			case MODIFIED:
				
				err := core.CopyFile(path, destPath)
				
				if err != nil {
					log.Println("error sync: ", err)
				}
			break

			case DELETED:
				err := os.RemoveAll(destPath)
				
				if err != nil {
					log.Println("error sync: ", err)
				}
				break

			default:
		}

		delete(w.ModFiles, path)
	}
}
