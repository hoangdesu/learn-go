package main

import "fmt"

func main() {
	fmt.Print("Enter your birth year: ")
	var birthYear int
	fmt.Scanf("%d", &birthYear)
	age := 2024 - birthYear
	fmt.Printf("You are %v years old.\n", age)

	isAdult := age >= 18
	
	if isAdult {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a kid")
	}

	// ---
	fmt.Print("Enter your money: ")
	var money int
	fmt.Scan(&money)

	mercedes := 100_000

	if money > mercedes {
		fmt.Println("You can buy a Mercedes 🎉")
	} else if money == mercedes {
		fmt.Println("Just enough to afford a Mercedes :/")
	} else {
		fmt.Println("Sorry u poor bitch")
	}
}