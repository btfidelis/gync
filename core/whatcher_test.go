package core 

import (
	"testing"
	"fmt"
)

func TestObserveFile(t *testing.T) {
	//ObserveFile("../t")
	ObserveDir("../t")
	fmt.Println("continuing with the program")
}