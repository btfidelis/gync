package core

import (
	"os"
	"fmt"
	"log"
)

/**
 * Observe file
 */
func ObserveFile(file string) {
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


func ObserveDir(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal("ON DAEMON: ", err)
	}

	infos, err := file.Readdir(5)

	if err != nil {
		log.Fatal("ON DAEMON: ", err)
	}

	for _, info := range(infos) {
		fmt.Println(info.Name());
	}
}







