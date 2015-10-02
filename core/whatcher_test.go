package core 

import (
	"testing"
	"fmt"
	"time"
)

func TestObserveFile(t *testing.T) {
	//ObserveFile("../t")
	var w Watcher
	w.ModTimes = make(map[string]time.Time, 0)
	
	w.ObserveDir("../t")
	
	fmt.Println("continuing with the program")
}