package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

//go:generate go-bindata -prefix "static/" -pkg main -o bindata.go static/...

const FILE_NOT_GIVEN = "You need to give a file name to convert to static site."

type DefaultValues struct {
	Input       string
	Path        string
	Title       string
	Author      string
	Description string
}

type Placeholders struct {
	Content     string
	Title       string
	Author      string
	Description string
}

type Config struct {
	Placeholder Placeholders
	Default     DefaultValues
	HelpText    DefaultValues
}

func main() {

	// Read and parse the config file.
	config := readConfig()

	// Define the available flags.
	inputFile := flag.String("input", config.Default.Input, config.HelpText.Input)
	resultDirectoryName := flag.String("path", config.Default.Path, config.HelpText.Path)
	projectTitle := flag.String("title", config.Default.Title, config.HelpText.Title)
	projectAuthor := flag.String("author", config.Default.Author, config.HelpText.Author)
	projectDescription := flag.String("description", config.Default.Description,
		config.HelpText.Description)

	flag.Parse()

	if *inputFile == config.Default.Input {
		pp(FILE_NOT_GIVEN)
		return
	}

	// Check if input file exists, and stop if not.
	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		pp("The file " + *inputFile + " does not exist.")
		return
	}

	// Read the template HTML file.
	template, err := Asset("template.html")
	check(err)

	// Read the markdown file.
	dat, err := ioutil.ReadFile(*inputFile)
	check(err)

	// Convert markdown to HTML and insert into the template.
	html := blackfriday.MarkdownCommon(dat)
	result := bytes.Replace(template, []byte(config.Placeholder.Content), html, -1)

	// Replace the placeholders with given values.
	result = bytes.Replace(result, []byte(config.Placeholder.Title), []byte(*projectTitle), -1)
	result = bytes.Replace(result, []byte(config.Placeholder.Author), []byte(*projectAuthor), -1)
	result = bytes.Replace(result, []byte(config.Placeholder.Description),
		[]byte(*projectDescription), -1)

	// Remove if an existing folder exists.
	if exists(*resultDirectoryName) {
		os.RemoveAll(*resultDirectoryName)
	}

	generateOutputFiles(*resultDirectoryName, result)

	pp("Documentation has been created successfully.\n" +
		"The sources can be found in '" + *resultDirectoryName + "' folder.")
}

// generateOutputFiles copies the required asset files to the specified location.
func generateOutputFiles(directoryPath string, data []byte) {

	// Create the result path and write HTML output to that path.
	os.Mkdir(directoryPath, 0644)
	os.Mkdir(directoryPath+"/css/", 0644)
	os.Mkdir(directoryPath+"/js/", 0644)
	err := ioutil.WriteFile(directoryPath+"/index.html", data, 0644)
	check(err)

	// Copy the asset files to the location.
	file, _ := Asset("css/default.min.css")
	os.Create(directoryPath + "/css/default.min.css")
	err = ioutil.WriteFile(directoryPath+"/css/default.min.css", file, 0644)
	check(err)

	file, _ = Asset("css/doc.css")
	os.Create(directoryPath + "/css/doc.css")
	err = ioutil.WriteFile(directoryPath+"/css/doc.css", file, 0644)
	check(err)

	file, _ = Asset("js/highlight.js")
	os.Create(directoryPath + "/js/highlight.js")
	err = ioutil.WriteFile(directoryPath+"/js/highlight.js", file, 0644)
	check(err)
}

// readConfig reads the config file from the static folder to a Config object.
func readConfig() (config Config) {
	configJson, _ := Asset("config.json")
	err := json.Unmarshal(configJson, &config)
	if err != nil {
		fmt.Println("Error occured while reading the config file:", err)
	}
	return config
}

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
func exists(path string) bool {
	found, err := os.Stat(path)
	if err == nil && found.IsDir() {
		return true
	}
	return false
}
