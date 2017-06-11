package main

import (
	"os"
	"encoding/json"
	"bytes"
    "io/ioutil"
    "io"
    "fmt"
    "path/filepath"
	"github.com/russross/blackfriday"
	"flag"
)

const FILE_NOT_GIVEN = "You need to give a file name to convert to static site."

type DefaultValues struct {
	Input string
	Path string
	Title string
	Author string
	Description string
}

type Placeholders struct {
	Content string
	Title string
	Author string
	Description string
}

type Config struct {
	Placeholder Placeholders
	Default DefaultValues
	HelpText DefaultValues
}

func main() {

	// Read and parse the config file.
	config := readConfig()
	
	// Define the available flags.
	inputFile 			:= flag.String("input", config.Default.Input, config.HelpText.Input)
	resultDirectoryName := flag.String("path", config.Default.Path, config.HelpText.Path)
	projectTitle 		:= flag.String("title", config.Default.Title, config.HelpText.Title)
	projectAuthor 		:= flag.String("author", config.Default.Author, config.HelpText.Author)
	projectDescription 	:= flag.String("description", config.Default.Description, 
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
    template, err := ioutil.ReadFile("static/template.html")
	check(err)

    // Read the markdown file.
    dat, err := ioutil.ReadFile(*inputFile)
    check(err)

    // Convert markdown to HTML and insert into the template.
	html := blackfriday.MarkdownCommon(dat)
	result := bytes.Replace(template, []byte(config.Placeholder.Content), html, -1)

	// Replace the placeholders with given values.
	result = bytes.Replace(result, []byte(config.Placeholder.Title), []byte(*projectTitle), -1)
	result = bytes.Replace(result, []byte(config.Placeholder.Description), []byte(*projectDescription), -1)
	result = bytes.Replace(result, []byte(config.Placeholder.Author), []byte(*projectAuthor), -1)

	// Remove if an existing folder exists.
	if(exists(*resultDirectoryName)) {
		os.RemoveAll(*resultDirectoryName)
	} 

	// Create the result path and write HTML output to that path.
	os.Mkdir(*resultDirectoryName, 0644)
	err = ioutil.WriteFile(*resultDirectoryName + "/index.html", result, 0644)
    check(err)

    // Copy the static files to the result folder.
    copyDir("static/js", *resultDirectoryName +"/js")
    copyDir("static/css", *resultDirectoryName + "/css")

    pp("Documentation has been created successfully.\n" + 
    	"The sources can be found in '" + *resultDirectoryName + "' folder.")
}

// readConfig reads the config file from the static folder to 
func readConfig() (config Config) {
	file, _ := os.Open("static/config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
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
func exists(path string) (bool) {
    found, err := os.Stat(path)
    if (err == nil && found.IsDir()) {
    	return true 
    }
    return false
}

// copyFile copies the contents of the file named src to the file named by dst. 
// The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// copyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func copyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("Given source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("Destination already exists")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = copyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}