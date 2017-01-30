package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer res.Body.Close()
	// Create a file to persist the response
	file, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	// Use MutiWriter so we can write to stdout and a file on
	// the same write operation
	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, res.Body)
}
