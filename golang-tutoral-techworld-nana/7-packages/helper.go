package main

import "fmt"

// this function will be visible to all files that share the same package main
func hi() {
	fmt.Println("Hi from helper!")
}