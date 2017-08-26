package main

import (
	"fmt"
	"os"
	"encoding/json"
)

const CONFIG_READ_ERROR = "Error occured while reading the config file: "

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

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// readConfig reads the config file from the static folder to a Config object.
func readConfig() (config Config) {
	configJson, _ := Asset("config.json")
	err := json.Unmarshal(configJson, &config)
	if err != nil {
		fmt.Println(CONFIG_READ_ERROR, err)
	}
	return config
}