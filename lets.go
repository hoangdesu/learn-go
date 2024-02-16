package main

import "fmt" // "Format package"

func main() {
	fmt.Print("Let's Go!")

	var name = "Brian"
	var username string = "hoangdesu"

	fmt.Println("username", username)

	fmt.Println(name + " Nguyen")

	name = "Hoang"
	fmt.Println(name + " Nguyen")

	fmt.Println("hello", "world")

	const birthYear int = 1995;
	// birthYear = 1997 -> error

	var age = 2024 - birthYear

	fmt.Println("You are", age, "years old")


	// Data types
	var password string

	fmt.Println("it's a", password)

	password = "secret"
	fmt.Println("it's a", password)

	hobby := "coding"
	fmt.Println(hobby)


	// Formatted print
	fmt.Printf("The value of variable \"name\" is \"%v\", data type is %T\n", name, name)
	// - can use %v for value, %T for data type
	// - or can use C-style formatted for correct data type e.g. %s for string..., %d for int 
	fmt.Printf("Age: %d", age)


	// User input
	// fmt.Print("Enter new username: ")
	// fmt.Scan(&username)
	// fmt.Println("New username:", username)

}


