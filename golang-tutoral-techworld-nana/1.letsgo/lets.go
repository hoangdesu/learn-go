package main

import "fmt" // "Format package"

func main() {
	fmt.Print("Let's Go!")

	var name = "Brian" // type string is auto inferred
	var username string = "hoangdesu" // can specify var type

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

	hobby := "coding" // declare + assign
	fmt.Println(hobby)


	// Formatted print
	fmt.Printf("The value of variable \"name\" is \"%v\", data type is %T\n", name, name)
	// - placeholders: can use %v for value in default format, %T for data type
	// - or can use C-style formatted for correct data type e.g. %s for string..., %d for int
	fmt.Printf("Age: %d, data type: %T\n", age, age)

	// int types: int 8 -> 64
	// same for uint

	// Pointers
	fmt.Printf("address of password %v: %p\n", password, &password)
	// - using %v or %p for pointer is fine, as long as we get the address of the variable with & e.g. &name


	// User input
	fmt.Printf("Current username: %v. Enter new username: ", username)
	fmt.Scan(&username) // - using the pointer to username's address to "pass by reference"
	fmt.Println("New username:", username)


	// Ex: Book tickets
	totalTickets := 10

	var tickets uint8

	fmt.Print("How many tickets: ")
	fmt.Scan(&tickets)

	if int(tickets) > totalTickets {
		fmt.Println("Error: Not enough tickets")
	} else {
		totalTickets -= int(tickets)
		fmt.Printf("User %v booked %v tickets. %v tickets remaining\n", username, tickets, totalTickets)
	}
}

// go run lets.go
// go run .
