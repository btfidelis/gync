package core

import(
	"encoding/json"
	"bufio"
	"os"
	"log"
)

const storagePath string = "./../storage"

type IOManager struct {
	Path string
}

func NewIOManager(path string) *IOManager {
	return &IOManager{path}
}

func (ioMan *IOManager) SaveObj(obj interface{}) {
	encoded, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(storagePath + ioMan.Path, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		log.Fatal(err)
	}

	if _, err = f.Write(encoded); err != nil {
		log.Fatal(err)
	}

	f.Sync()
	f.Close()
}


func (ioMan *IOManager) LoadFile(obj interface{}) []interface{} {
	file, err := os.OpenFile(storagePath + ioMan.Path, os.O_RDONLY, 0600)
	objectCol := make([]interface{}, 0)


	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err = json.Unmarshal(scanner.Bytes(), obj)

		if err != nil {
			log.Fatal(err)
		}

		objectCol = append(objectCol, obj)
	}

	return objectCol
}
