package model 

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"errors"
	"regexp"
	"encoding/json"
	"github.com/btfidelis/gync/core"
)

type SaveCollection struct {
	Saves []Save
}

type Save struct{
	Name		string
	Location 	string
	Dir			bool
}

func NewSave(name string, local string) *Save {
	errName := validateName(name)
	file, errPath := os.Stat(local)

	if errName != nil {
		fmt.Println(errName)
		return nil
	}

	if errPath != nil {
		fmt.Println("The path provided does not exist")
		return nil
	}

	return &Save{name, local, file.IsDir()}
}

func GetSaveLocal() string {
	_, file, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(file), "../storage/saves.json")
}

/**
 * Validates if the name is unique and valid
 */
func validateName(name string) error {
	regex, err := regexp.Compile("(^[\\w\\-]+$)")

	if err != nil {
		log.Fatal("invalid regex")
	}

	if !validateUniqueName(name) {
		return errors.New("You must enter a unique name (gync list to see used names)")
	}

	matched := regex.MatchString(name)

	if !matched {
		return errors.New("You must enter a valid name (Only alphanumeric and \"-\" symbol no spaces)")
	}

	return nil
}

/**
 *  Returns true if name is unique
 */
func validateUniqueName(name string) bool {
	saveCol := GetSaveCollection()

	for _, col := range(saveCol.Saves) {
		if col.Name == name {
			return false
		}
	}

	return true
}


func (save *Save) Save() {
	io := core.NewIOManager(GetSaveLocal())
	saveCol := GetSaveCollection()

	saveCol.Saves = append(saveCol.Saves, *save)

	saves, err := json.Marshal(saveCol)

	if err != nil {
		log.Fatal(err)
	}

	io.SaveObj(saves)
}

func GetSaveCollection() SaveCollection {
	io := core.NewIOManager(GetSaveLocal())
	var saveCol SaveCollection
	saves := io.LoadFile()

	err := json.Unmarshal(saves, &saveCol)

	if err != nil {
		log.Fatal(err)
	}

	return saveCol
}

func (saveCol SaveCollection) Where(name string) (*Save, int) {
	
	for i, save := range(saveCol.Saves) {
		if (save.Name == name) {
			return &save, i
		}
	}

	return nil, -1
}

func (saveCol *SaveCollection) Remove(id int) {
	io := core.NewIOManager(GetSaveLocal())

	if id != -1 {
		saveCol.Saves = append(saveCol.Saves[:id], saveCol.Saves[id+1:]...)
	}

	saves, err := json.Marshal(saveCol)

	if err != nil {
		log.Fatal("On Delete:", err)
	}

	io.SaveObj(saves)
}
