package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// There are various ways to declare a map in Go
	// Method 1: Using var keyword
	var m1 map[string]string
	fmt.Println(m1) // nil

	// Method 2: Using make function
	m2 := make(map[string]string)
	fmt.Println(m2) // map[]

	// Method 3: Using literal syntax
	m3 := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
	}
	fmt.Println(m3) // map[one:1 two:2 three:3]

	// Need to use [] syntax with maps as the type can be other things than strings
	fmt.Println(m3["one"])   // 1
	fmt.Println(m3["two"])   // 2
	fmt.Println(m3["three"]) // 3

	delete(m3, "one")
	fmt.Println(m3) // map[two:2 three:3]

	for key, value := range m3 {
		fmt.Println(key, value)
	}
}
