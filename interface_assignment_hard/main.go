package main

import (
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	// Defer schedules the file to close just before the main function exits so that resources are properly disposed of
	defer file.Close()

	io.Copy(os.Stdout, file)
}
