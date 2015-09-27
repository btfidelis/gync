package core

import(
	"io/ioutil"
	"os"
	"log"
)

const storagePath string = "storage"

type IOManager struct {
	Path string
}

func NewIOManager(path string) *IOManager {
	return &IOManager{path}
}

func (ioMan *IOManager) SaveObj(obj []byte) {
	
	f, err := os.OpenFile(storagePath + ioMan.Path, os.O_WRONLY, 0600)

	if err != nil {
		log.Fatal(err)
	}

	if _, err = f.Write(obj); err != nil {
		log.Fatal(err)
	}

	f.Sync()
	f.Close()
}


func (ioMan *IOManager) LoadFile() []byte {

	b, err := ioutil.ReadFile(storagePath + ioMan.Path)
	
	if err != nil {
		log.Fatal("Unable to open file")
		return nil
	}

	return b
}