package main

// - acts as import path for all packages
// - file name doesn't matter, import with package's name (same as folder's name)
// - if simply import "utils" custom package, Go will look for the built-in module -> need to specify the path

import (
	"fmt"
	"learn-go/7-packages/utils" // module learn-go comes from go.mod
)

func main() {
	hi()                     // coming from helper file, same package -> no need to import
	utils.Introduce("Hoàng") // import from package utils

	// can also import package-level var and const
	fmt.Printf("Nickname %v, born in %v\n", utils.Nickname, utils.Birthyear)

	// can directly re-assign value in other package:
	utils.Nickname = "Brian"
	fmt.Println("New nickname:", utils.Nickname)
}
