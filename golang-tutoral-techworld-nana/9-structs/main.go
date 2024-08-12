package main

import "fmt"

// structs: mixed data types
// - lightweight class -> doesn't support inheritance

type ProgrammingLanguage struct {
	name        string // fields are separated by new lines, no need comma or semicolon
	year        uint
	isLoved     bool
	developedBy []string
}

// type alias
type ProgrammingLanguages = []ProgrammingLanguage

func main() {
	// create an array of struct
	programmingLanguages := make(ProgrammingLanguages, 0) // same as typing []ProgrammingLanguage

	python := ProgrammingLanguage{
		name:        "Python",
		year:        1991,
		isLoved:     true,
		developedBy: []string{"Guido van Rossum"},
	}

	programmingLanguages = append(programmingLanguages, python)

	programmingLanguages = append(programmingLanguages, ProgrammingLanguage{
		name:        "Go lang",
		year:        2009,
		isLoved:     true,
		developedBy: []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"},
	})

	fmt.Println(programmingLanguages)

	// can re-assign fields
	programmingLanguages[1].name = "Go"


	// access properties with dot notation
	fmt.Printf("Python: %v\n", python.year)

	for _, lang := range programmingLanguages {
		fmt.Println(lang.format())
	}

}

// receiver function
// - similar to methods in a class, can only be called on an object
// - invoke with object.method()
// - pass in a pointer to the struct, otherwise it will pass by value (value is copied)
func (lang *ProgrammingLanguage) format() string {
	// (*lang).name // another syntax to deference the struct object
	// however, we can access directly without having to deference (done by go by default)

	var s string

	s += fmt.Sprintf("- %v: developed in %v by ", lang.name, lang.year)

	for i, dev := range lang.developedBy {
		s += fmt.Sprintf("%v", dev)
		if i < len(lang.developedBy)-1 {
			s += ", "
		}
	}

	return s
}
