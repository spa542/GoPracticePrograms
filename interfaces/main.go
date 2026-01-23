package main

import "fmt"

// Interface definition
type bot interface {
	getGreeting() string
	//getGreeting(int, string) (string, error)
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// Without Interfaces

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// With Interfaces

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Custom logic for generating an english greeting
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
