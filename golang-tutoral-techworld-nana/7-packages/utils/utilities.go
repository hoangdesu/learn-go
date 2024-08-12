package utils

import "fmt"

// exported function: capitalize the first letter
func Introduce(name string) {
	fmt.Printf("Hi guys my name is %v\n", name)
}

// variables and constants can also be exported, with capitalized first letter
var Nickname string = "Cún"
const Birthyear = 1995

// these vars and funcs can be seen as global scope