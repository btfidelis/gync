package app 

import (
	"testing"
	"fmt"
	"time"
)

func TestObserveFile(t *testing.T) {
	//ObserveFile("../t")
	w := Watcher{
		ModTimes: make(map[string]time.Time, 0),
		ModFiles: make(map[string]uint32, 0),
	}

	w.ObserveDir("../t")
	
	fmt.Println("continuing with the program")
}