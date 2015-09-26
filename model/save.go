package model 

import (
	"fmt"
	"regexp/syntax"
)

type Save struct{
	Name		string
	Location 	string
	Dir			bool
}


func (save Save) NewSave(name string, local string) *Save {

}


func validateName(name string) bool {
	matched, err := syntax.MatchString("/[A-Z|a-z|0-9|\\-]/", name)

	if err != nil {
		return matched

	} else {
		fmt.Println("You must expecify a name")
	}
}


