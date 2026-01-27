package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	// This is just to show that there is a function that can read the exact number of bytes from the response body
	// body, err := io.ReadAll(resp.Body)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	lw := logWriter{}

	// Writer interface
	// The new reader is required because the response body has already been read from the buffer. You can reuse the buffer with "bs" variable if you create a new reader
	io.Copy(os.Stdout, bytes.NewReader(bs))
	// Another way!
	io.Copy(lw, bytes.NewReader(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this number of bytes:", len(bs))
	return len(bs), nil
}
