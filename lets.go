package main

import "fmt"

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
	password = "secret"

	fmt.Println("it's a", password)

	hobby := "coding"
	fmt.Println(hobby)


	// User input
	fmt.Print("Enter new username: ")
	fmt.Scan(&username)
	fmt.Println("New username:", username)
}
