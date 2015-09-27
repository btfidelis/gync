package core

import(
	"path/filepath"
	"io/ioutil"
	"os"
	"log"
	"fmt"
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
	
	fmt.Println(dir)
	return dir
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
	b, err := ioutil.ReadFile(ioMan.GetPath())
	
	if err != nil {
		log.Fatal("Unable to open file")
		return nil
	}

	return b
}