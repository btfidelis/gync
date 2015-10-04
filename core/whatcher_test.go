package core 

import (
	"testing"
	"fmt"
	"time"
)

func TestObserveFile(t *testing.T) {
	//ObserveFile("../t")
	w := Watcher{ModTimes: make(map[string]time.Time, 0)}

	w.ObserveDir("../t")
	
	fmt.Println("continuing with the program")
}