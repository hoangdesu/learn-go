package main

import "fmt"

func main() {
	// Arrays
	// arrays in Go are fixed size

	var programmingLanguages [50]string // declare an empty string array of size 50

	// can also use:
	// 	var programmingLanguages = [50]string{} // init empty array
	// or put initial elements:
	// 	var programmingLanguages = [50]string{"Python", "Go"} 


	// Add new elements:
	programmingLanguages[0] = "JavaScript"
	programmingLanguages[1] = "Python"
	programmingLanguages[2] = "Go"
	programmingLanguages[10] = "Swift"

	fmt.Printf("All languages: %v\n", programmingLanguages)
	fmt.Printf("Learning: %v\n", programmingLanguages[2])
	fmt.Printf("Array type: %T\n", programmingLanguages)
	fmt.Printf("Array length: %v\n", len(programmingLanguages))

	// Slices
	// - an abstraction of arrays
	// - uses array under the hood, but has dynamic size
	// same syntax as array, just dont define size :/
	var games []string 
	// - can also use syntax: 
		// - var games = []string{} 
		// - games := []string{}

	games = append(games, "League of Legends") // add new element at the end, not as clean as python :((
	games = append(games, "Overwatch")

	fmt.Printf("games: %v\n", games)
	fmt.Printf("slice type: %T\n", games)
	fmt.Printf("slice length: %v\n", len(games))
	fmt.Printf("index 0: %v\n", games[0])

}