package model

import(
	"testing"
)


func TestValidateName(t *testing.T) {
	cases := []struct {
		in string
		want bool 
	} {
		{"Dark Souls", true}, 
		{"dark-souls", true}, 
		{"darksouls2", true}, 
		{"half-life3", true},
		{".. praise it", false},
		{"!raise!your!dongers", false},
	}

	for _, c := range(cases) {
		got := model.validateName(c.in)
		if got != c.want {
			t.Errorf("Error asserting that %s is %b", c.in, c.want)
		}
	}
}