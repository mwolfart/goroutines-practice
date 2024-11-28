package io

import "fmt"

func ParseInput(args []string) int {
	if (len(args) != 8) {
		fmt.Println("Invalid number of arguments")
		return 1
	}
	return 0
}