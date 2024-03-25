package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	// mockFetchData("pizza") // this blocks the code for 3 seconds before we can continue
	// => should be executed in a different thread

	// use "go" to execute a function in a gorountine (lightweight thread)
	// go mockFetchData("pizza")

	// however the execution above will finish any time in future
	// by default, the main goroutine will NOT wait for other goroutines

	// we need to tell the main thread to "wait" for other threads to complete before it can exit
	// => use sync.WaitGroup()

	// set the number of goroutines to wait for
	// add before creating a goroutine
	wg.Add(1)
	go mockFetchData("pizza")

	// go must be called before a function -> create an annonymous function and immediately invoke 
	// similar to IIFE in JS
	wg.Add(1)
	go func() {
		var food string
		fmt.Print("What food: ")
		fmt.Scan(&food)
		mockFetchData(food)
		// wg.Done() // adding Done() here somehow gets error "panic: sync: negative WaitGroup counter" :/
	}() 

	fmt.Println("This code is executed in the main thread")

	// blocks the main thread from exiting until the WaitGroup counter is 0
	wg.Wait()

}

func mockFetchData(food string) {
	fmt.Printf("Fetching data for %v...\n", food)
	time.Sleep(3 * time.Second) // block the current thread
	fmt.Println("\nDone fetching", food)

	// call this function at the end of each goroutine, decrease the counter by 1
	wg.Done() 
}
