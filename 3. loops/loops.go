package main

import "fmt"

func main() {
	// - loops are simplified in Go
	// - only have for loop, no while or for-each

	// - a for loop without condition is a while loop
	// infiniteLoop()

	growingArray()


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

		var newFood string
		fmt.Print("\nEnter a new food: ")
		fmt.Scan(&newFood)

		if newFood == "bye" {
			fmt.Println("Bye!")
			break
		} else {
			foods = append(foods, newFood)
		}
	}
}
