package main

import (
	"fmt"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// pp prints the given string, built to make code more readable.
func pp(s string) {
	fmt.Println(s)
}

// exists returns whether the given file or directory exists or not
func dirExists(path string) bool {
	found, err := os.Stat(path)
	if err == nil && found.IsDir() {
		return true
	}
	return false
}
