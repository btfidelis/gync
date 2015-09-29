package core

import(
	"path/filepath"
	"io/ioutil"
	"log"
)

const storagePath string = "../storage/"

type IOManager struct {
	Path string
}

func NewIOManager(path string) *IOManager {
	return &IOManager{path}
}

func (io IOManager) GetPath() string {
	dir, _ := filepath.Abs(storagePath + io.Path)
	
	return dir
}

func (ioMan *IOManager) SaveObj(obj []byte) {
	
	err := ioutil.WriteFile(storagePath + ioMan.Path, obj, 0600)

	if err != nil {
		log.Fatal(err)
	}
}


func (ioMan *IOManager) LoadFile() []byte {
	b, err := ioutil.ReadFile(ioMan.GetPath())
	
	if err != nil {
		log.Fatal("Unable to open file")
		return nil
	}

	return b
}