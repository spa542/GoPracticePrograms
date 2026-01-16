package main

import "fmt"

func main() {

	numList := make([]int, 11) // 0 to 10 inclusive
	for i := range numList {
		numList[i] = i
	}

	for _, num := range numList {
		if num%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
