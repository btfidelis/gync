package core

import "testing"
import "fmt"
import "encoding/json"

type user struct {
	Name string
	Nickname string
}

type collection struct {
	Elements []*user
}

func TestIOManWrite(t *testing.T) {
	io := NewIOManager("/testing.json")

	u := &user{"Hey", "WuckFaffle"}
	u1 := &user{"BLHA", "FISH"}

	var col collection
	col.Elements = append(col.Elements, u, u1)


	encoded, err := json.Marshal(col)

	if err != nil {
		t.Errorf("Encoding error")
	}

	io.SaveObj(encoded)
}

func TestLoadFile(t *testing.T) {
	io := NewIOManager("/testing.json")

	var col collection;
	colBytes := io.LoadFile() 

	if colBytes == nil {
		return
	}
	err := json.Unmarshal(colBytes, &col)
	
	if err != nil {
		t.Errorf("error decoding: ", err)
	}

	fmt.Println(*col.Elements[0])
}

func TestCopyFile(t *testing.T) {
	err := CopyFile("../t/copy_test.txt", "../t/t/copiatest2.txt")

	if err != nil {
		t.Errorf("error: ", err)
	}
}