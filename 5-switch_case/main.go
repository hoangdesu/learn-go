package main

import "fmt"

func main() {
	var country string
	fmt.Print("Where are you from: ")
	fmt.Scan(&country)

	var greeting string
	switch country {
	case "vietnam":
		greeting = "xin chào"
	case "japan":
		greeting = "konnichiwa"
	case "france":
		greeting = "bonjour"
	default:
		greeting = "hello"
	}

	fmt.Printf("%v!\n", greeting)

}

// no need to use "break"
