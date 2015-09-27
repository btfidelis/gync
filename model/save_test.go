package model

import(
	"testing"
)


func TestValidateName(t *testing.T) {

	//errorMsg := errors.New("You must enter a valid name (Only alphanumeric and \"-\" symbol)")

	cases := []struct {
		in string
		want error 
	} {
		//{"Dark Souls", errorMsg}, 
		{"dark-souls", nil}, 
		{"darksouls2", nil}, 
		{"word", nil}, 
		{"half-life3", nil},
		//{".. praise it", errorMsg},
		//{"!raise!your!dongers", errorMsg},
	}

	for _, c := range(cases) {
		got := validateName(c.in)
		if got != c.want {
			t.Errorf("Error asserting that %s is ", c.in, c.want)
			t.Errorf("\nGot: ", got)

		}
	}
}


func TestNewSaveObject(t *testing.T) {
	cases := []struct {
		name, local string
	} {
		{
			"dark-souls",
		 	"C:\\Program Files (x86)\\Steam\\steamapps\\common\\Dark Souls Prepare to Die Edition",
		},
		{
			"something",
		 	"C:\\Program Files (x86)\\Steam\\steamapps\\common\\Dark Souls Prepare to Die Edition",
		},
	}


	for _, c := range(cases) {
		got := NewSave(c.name, c.local)
		if got == nil {
			t.Errorf("Error asserting that is Obj, ", got)
		}
	}
}


func TestNewSaveObjectPersist(t *testing.T) {
	save := NewSave("ShadowOfMordor", "C:\\Program Files (x86)\\Steam\\steamapps\\common\\")
	save1 := NewSave("Diablo2", "C:\\Games")

	if save == nil || save1 == nil {
		t.Errorf("Error creating obj")
		return
	}

	save.Save()
	save1.Save()
}


func TestValidateUniqueName(t *testing.T) {
	cases := []struct{
		name string
		assert bool
	}{
		{"Diablo2", false},
		{"testing", true},
	}

	for _, c := range(cases) {
		got := validateUniqueName(c.name)

		if got != c.assert {
			t.Errorf("Failed asserting that ", c.assert, "is ", got)
		}
	}
}

func TestRemoval(t *testing.T) {
	saves := GetSaveCollection()
	save, id := saves.Where("ShadowOfMordor")

	if save != nil {
		saves.Remove(id)
	} else {
		t.Errorf("Failed deleting, not found")
	}
}