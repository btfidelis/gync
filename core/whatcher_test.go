package core 

import (
	"testing"
	"fmt"
)

func TestObserveFile(t *testing.T) {
	//ObserveFile("../t")
	var w Watcher
	
	w.ObserveDir("../t")
	
	fmt.Println("continuing with the program")
}