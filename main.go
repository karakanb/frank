package main

import (
	"os"
	"bytes"
    "io/ioutil"
    "io"
    "fmt"
    "path/filepath"
	"github.com/russross/blackfriday"
	"flag"
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

func main() {

	// Config placeholders.
	templatePlaceholder := []byte("$$CONTENT$$")
	titlePlaceholder := []byte("$$TITLE$$")
	authorPlaceholder := []byte("$$AUTHOR$$")
	descriptionPlaceholder := []byte("$$DESCRIPTION$$")

	// Default values.
	defaultTitle := "Docs"
	defaultDescription := "Documentation of an awesome project."
	defaultAuthor := "Project Author"

	// Define the available flags.
	resultDirectoryName := flag.String("path", "result", "The path to export the resulting files.")
	projectTitle := flag.String("title", defaultTitle, "Title of the documentation page.")
	projectDescription := flag.String("description", defaultDescription, 
		"Description of the project to place in the 'description' meta tag.")
	projectAuthor := flag.String("author", defaultAuthor,
		"Author of the project to place in the 'author' meta tag.")
	
	flag.Parse()

    // Read the template HTML file.
    template, err := ioutil.ReadFile("static/template.html")
	check(err)
    
    // Read the markdown file.
    dat, err := ioutil.ReadFile("input.md")
    check(err)

    // Convert markdown to HTML and insert into the template.
	html := blackfriday.MarkdownCommon(dat)
	result := bytes.Replace(template, templatePlaceholder, html, -1)

	// Replace the placeholders with given values.
	result = bytes.Replace(result, titlePlaceholder, []byte(*projectTitle), -1)
	result = bytes.Replace(result, descriptionPlaceholder, []byte(*projectDescription), -1)
	result = bytes.Replace(result, authorPlaceholder, []byte(*projectAuthor), -1)
	
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