package main

import (
	"os"
	"bytes"
    "io/ioutil"
	"github.com/russross/blackfriday"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	templatePlaceholder := []byte("$$CONTENT$$")

    // Read the template HTML file.
    template, err := ioutil.ReadFile("static/template.html")
	check(err)
    
    // Read the markdown file.
    dat, err := ioutil.ReadFile("input.md")
    check(err)

	html := blackfriday.MarkdownCommon(dat)
	result := bytes.Replace(template, templatePlaceholder, html, -1)

	os.MkDir("result", 0644)
	err = ioutil.WriteFile("index.html", result, 0644)
    check(err)

}