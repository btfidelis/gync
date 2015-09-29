package core

import (
	"os"
	"fmt"
	"log"
	"time"
	"path/filepath"
)

type Watcher struct {
	ModTimes []time.Time
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

	w.ModTimes = append(w.ModTimes, info.ModTime())

	return err
}

func (m *mirror) WalkMirror(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(w.ModTimes[*w.Step])

	return err	
}

func (w *Watcher) ObserveDir(path string) {
	err := filepath.Walk(path, w.WalkPopulate)
	mirror := new(Watcher) 

	if err != nil {
		log.Fatal("ON DAEMON: ", err)
	}

	copy(mirror.ModTimes, w.ModTimes)

	for {
		err := filepath.Walk(path, mirror.WalkPopulate)
		
		if err != nil {
			log.Fatal("ON DAEMON: ", err)
		}

		if w.check(mirror) {
			copy(w.ModTimes, mirror.ModTimes)
		}
	}
}

func (w *Watcher) check() bool {

}


