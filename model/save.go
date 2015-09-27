package model 

import (
	"fmt"
	"log"
	"os"
	"errors"
	"regexp"
	"github.com/btfidelis/gync/core"
)

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
		fmt.Println(errPath)
		return nil
	}

	return &Save{name, local, file.IsDir()}
}


func validateName(name string) error {
	regex, err := regexp.Compile("(^[\\w\\-]+$)")

	if err != nil {
		log.Fatal("invalid regex")
	}

	matched := regex.MatchString(name)

	if !matched {
		return errors.New("You must enter a valid name (Only alphanumeric and \"-\" symbol no spaces)")
	} 

	return nil
}


func (save *Save) Save() {
	io := core.NewIOManager("/saves.json")
	io.SaveObj(save)
}

