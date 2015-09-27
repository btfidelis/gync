package core

import "testing"
import "fmt"

type user struct {
	Name string
	Nickname string
}

func TestIOManWrite(t *testing.T) {
/*	io := NewIOManager("/testing.json")

	u := &user{"Hey", "WuckFaffle"}
	u1 := &user{"ghost", "butts"}

	io.SaveObj(u)
*/}

func TestLoadFile(t *testing.T) {
	io := NewIOManager("/testing.json")

	var u user;

	userCol := io.LoadFile(&u)

	fmt.Println(user(userCol[0]))

}