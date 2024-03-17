package main

import (
	"fmt"
	"strings"
)

func main() {
	// - loops are simplified in Go
	// - only have for loop, no while or for-each

	// - a for loop without condition is a while loop
	// infiniteLoop()

	// growingArray()

	// split string with strings.Fields(), import from "strings" pkg
	// - white space as separator
	// - return a slice

	// games := strings.Fields("lienminh overwatch")
	// fmt.Printf("\nGames: %v\n", games)

	// blank identifier: _ (underscore)
	// 	- ignore a var

	// whileLoop()

	validateInput()
}

func infiniteLoop() {
	name := "Hoang"
	for {
		fmt.Printf("old name: %s. Enter new name: ", name)
		fmt.Scan(&name)
		if name == "quit" {
			break
		} else {
			fmt.Printf("new name: %v\n", name)
		}
	}
}

func growingArray() {
	foods := []string{}

	for {
		// fmt.Printf("\nFood slices: %v\n", food)

		// using the "range" keyword
		for index, food := range foods {
			fmt.Printf("index: %v, food: %v\n", index, food)
		}

		// range returns 2 values: index, sliceElement

		var newFood string
		fmt.Print("\nEnter a new food: ")
		fmt.Scan(&newFood)

		if newFood == "bye" {
			fmt.Println("Bye!")
			fmt.Printf("Final food slice: %v", foods)
			break
		} else {
			foods = append(foods, newFood)
		}
	}
}

func whileLoop() {
	// - for loop with condition == while loop
	i := 1
	for i < 10 {
		fmt.Println(i);
		i++;
	}
}

func validateInput() {
	for {
		fmt.Print("Enter email address: ")
		var email string
		fmt.Scan(&email)

		// check if string is missing character @
		if !strings.Contains(email, "@") {
			fmt.Println("Invalid email format")
			continue
		} else {
			fmt.Printf("Welcome %v!\n", email)
			break
		}
	}
}

// - continue: skip the remaining code
// - break: quit the loop
// && AND
// || OR
// ! NOT
// for true == while true