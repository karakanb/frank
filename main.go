package main

import (
	"bytes"
	"flag"
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

	// Check if input file is given as a parameter.
	if *inputFile == config.Default.Input {
		pp(FILE_NOT_GIVEN)
		return
	}

	// Check if input file exists, and stop if not.
	if !fileExists(*inputFile) {
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
	if dirExists(*resultDirectoryName) {
		os.RemoveAll(*resultDirectoryName)
	}

	generateOutputFiles(*resultDirectoryName, result)

	pp("Documentation has been created successfully.\n" +
		"The sources can be found in '" + *resultDirectoryName + "' folder.")
}

// generateOutputFiles copies the required asset files to the specified location.
func generateOutputFiles(directoryPath string, data []byte) {

	// Create the result path and write HTML output to that path.
	os.Mkdir(directoryPath, 0755)
	os.Mkdir(directoryPath+"/css/", 0755)
	os.Mkdir(directoryPath+"/js/", 0755)
	err := ioutil.WriteFile(directoryPath+"/index.html", data, 0755)
	check(err)

	// Copy the asset files to the location.
	fileNames := [...]string{"css/default.min.css", "css/doc.css", "js/highlight.js"}
	for _, path := range fileNames {
		file, _ := Asset(path)
		os.Create(directoryPath + "/" + path)
		err = ioutil.WriteFile(directoryPath+"/"+path, file, 0755)
		check(err)
	}
}
