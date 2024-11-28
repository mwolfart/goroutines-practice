package main

import (
	"fmt"
	"os"
	"mwolfart/goroutines-practice/src/io"
)

func main() {
	args := os.Args
	code := io.ParseInput(args)
	if code != 0 {
		os.Exit(code)
	}
	fmt.Println("All done!")
}