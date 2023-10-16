package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// when we run "go run main.go ../notes.txt" below command `os.Args[1]` takes first index which is ../notes.txt from the runned command
	// fmt.Println(os.Args[1])
	// os.Open("../notes.txt")
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// This will read and display the file at first index
	io.Copy(os.Stdout, f)
}

// Tip to read the main.go file do "go build main.go" and then run "./main main.go"
