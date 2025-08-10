package main

import "fmt"

func main() {
	// Creating a Map in Go
    var ages map[string]int
    
    // Initializing the Map
    ages = make(map[string]int)

    // Adding Values
    ages["Alice"] = 30
    ages["Bob"] = 25

    fmt.Println(ages["Alice"]) // Output: 30
}