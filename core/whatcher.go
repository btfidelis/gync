package core

import (
	"os"
	"fmt"
	"log"
	"time"
	"path/filepath"
)

type Watcher struct {
	ModTimes 	[]time.Time
	Name		string
}

/**
 * Observe file
 */
func (w Watcher) ObserveFile(file string) {
	info, err := os.Stat(file)

	if err == nil {
		lastChangedOn := info.ModTime()

		for {
		
			info, err := os.Stat(file)
			
			if err == nil {

				if info.ModTime().After(lastChangedOn) {
					fmt.Println("arquivo alterado: ", info.Name())
					lastChangedOn = info.ModTime()
				}
			}
		}
	}
}

func (w *Watcher) WalkPopulate(path string, info os.FileInfo, err error) error {

	if err != nil {
		log.Fatal(err)
		return err
	}

	w.ModTimes = make([]time.Time, 0)
	w.ModTimes = append(w.ModTimes, info.ModTime())

	return err
}

func (w *Watcher) ObserveDir(path string) {
	err := filepath.Walk(path, w.WalkPopulate)
	mirror := new(Watcher) 
	
	if err != nil {
		log.Fatal("ON DAEMON: ", err)
		return
	}

	mirror.ModTimes = make([]time.Time, len(w.ModTimes))

	copy(mirror.ModTimes, w.ModTimes)
	
	for {
		err := filepath.Walk(path, mirror.WalkPopulate)
		
		if err != nil {
			log.Fatal("ON DAEMON: ", err)
			return
		}

		if w.check(mirror.ModTimes) {
			fmt.Println("alterado")
			copy(w.ModTimes, mirror.ModTimes)		
		}

		mirror.ModTimes = make([]time.Time, len(w.ModTimes), len(w.ModTimes))
	}
}

func (w *Watcher) check(mirror []time.Time) bool {
	if len(w.ModTimes) != len(mirror) {
		fmt.Println(len(w.ModTimes)," != ", len(mirror))
		return true
	}

	for i := 0; i < len(mirror); i++ {
		if ! w.ModTimes[i].Equal(mirror[i]) {
			fmt.Println(w.ModTimes[i]," != ", mirror[i])
			return true
		}
	}

	return false
}


