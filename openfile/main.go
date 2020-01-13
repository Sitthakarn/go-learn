package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	readFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	dataRead, err := os.Open(readFile.Name())
	if err != nil {
		fmt.Println("Error:", err)
	}

	io.Copy(os.Stdout, dataRead)
}
