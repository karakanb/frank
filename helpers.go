package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
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
func pp(s interface{}) {
	fmt.Println(s)
}

// remove removes an element from a given array.
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// dirExists returns whether the given file or directory exists or not
func dirExists(path string) bool {
	found, err := os.Stat(path)
	if err == nil && found.IsDir() {
		return true
	}
	return false
}

// fileExists returns whether the given file exists or not
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
		printError(CONFIG_READ_ERROR + err.Error())
	}
	return config
}

func copyFile(directoryPath string, filePath string) {
	file, _ := Asset(filePath)
	os.Create(directoryPath + "/" + filePath)
	err := ioutil.WriteFile(directoryPath+"/"+filePath, file, 0755)
	check(err)
}

func printSuccess(s string) {
	color.Green(wrapStars(s))
}

func printError(s string) {
	color.Red(wrapStars(s))
}

func wrapStars(s string) string {
	stars := "\n************************************************************\n"
	s = stars + "\n" + s + "\n" + stars
	return s
}
