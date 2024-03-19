package main

import (
	"fmt"
	"strconv" // string conversions
)

func main() {
	// Map: unique key-value pairs
	// - map[key-type]value-type -> only define map data type
	// - use "make()" function: make(map[key-type]val-type) -> init an empty map
	// - all keys have the same data type; all values have the same data type
	// - => no mixed values :(

	zed := make(map[string]string)
	zed["name"] = "Zed"
	zed["R"] = "Death Mark"
	zed["energy"] = "200" // cannot use int here
	zed["health"] = strconv.FormatUint(uint64(654), 10) // first convert to uint65, specify base 10 (decimal), then convert to string

	fmt.Println("zed map:", zed) // -> map[Q:Razor Shuriken R:Death Mark energy:200 health:654]


	// declare and init a map, no need to use make()
	ahri := map[string]string{"name": "Ahri", "R": "Spirit Rush", "mana": fmt.Sprint(418), "health": fmt.Sprintf("%v", 590) }
	// btw cannot convert int to string with string(418) -> yields a rune, incorrect desire
	// apart from strconv, can also use fmt.Sprint(num) or fmt.Sprintf to convert to string


	// Slice of Maps
	// can also use make() to create an empty slice
	champions := make([]map[string]string, 0) // reads: 'make an array of type map, with key as string and value as string". Init to 0 element 
	
	champions = append(champions, zed, ahri) // can append multiple elements at once. First param is a slice

	fmt.Println("All champions:", champions) // -> All champions: [map[R:Death Mark energy:200 health:654 name:Zed] map[R:Spirit Rush health:590 mana:418 name:Ahri]]

	fmt.Println("Their ultimates:")
	for i, champ := range champions {
		fmt.Printf("%v. %v's R: %v\n", (i+1), champ["name"], champ["R"])
	}
}

