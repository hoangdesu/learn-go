package main

import "fmt"

// package-level variables:
//	- shared among functions
// 	- cannot be created with the := syntax, must use var to declare
// 	- not a good practice. Try to keep variables as local as possible

var myName = "Cún"

func main() {

	greetUser("Brian")
	introduce("Hoang", 29, true)

	doubleArray := doubleUp([]int{6, 9, 4, 2})
	fmt.Printf("Double array: %v\n", doubleArray)

	// store return values in order of return
	name, lane, level, isMeta := getZedInfo()
	fmt.Printf("Champ: %v - Lane: %v - Level: %v - Is meta: %v\n", name, lane, level, isMeta)
	
	// package-level var
	whoAmI()
	changeMyName() // can accidentally be modified by other funcs
	whoAmI() // myName has been changed
	

	// variadic functions
	sumAllNums(1, 2, 3)

	// can "spread" (...) all elements in a slice to use as args -> e.g. func(slice...)
	nums := []int{9, 1, 2, 3, 4, 6}
	sumAllNums(nums...)
	
}

// function is hoisted, can be defined after main
func greetUser(name string) {
	fmt.Printf("Hello %v\n", name)
}


// multiple params
func introduce(name string, age int, handsome bool) {
	if handsome {
		fmt.Printf("Hi guys, my name is %v, I am %v years old and I am definitely handsome!\n", name, age)
	}
}


// return value, make sure to define return type
func doubleUp(nums []int) []int {
	doubleNums := []int{}

	for _, num := range nums {
		doubleNums = append(doubleNums, num*2)
	}

	return doubleNums
}


// return multiple values: must wrap all values inside a ()
// order of values matter
func getZedInfo() (string, string, uint, bool) {
	return "Zed", "mid", 18, true
}


func whoAmI() {
	fmt.Printf("My name is %v\n", myName) // package-level variable
}

func changeMyName() {
	myName = "Brian"
}


// Variadic Functions: can be called with any number of trailing arguments
// e.g. fmt.Println()
// take an arbitrary number of ints as arguments
func sumAllNums(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Total = %v\n", total)
}

