# Frank

Frank is a simple static site generator based on Markdown, with basic documentation needs in mind. It basically takes the markdown input, renders it and creates a sample output folder, which can be deployed to a very basic static site hosting service, such as GitHub Pages.

## Usage

In order to start using it directly, you can use `go get` to get the tool:
```bash
# Get the library and install it.
go get github.com/karakanb/frank

# Create static site from README.md file. Resulting static files will be in result/ folder.
frank
```

### Available Flags

Keep in mind that, all of these flags are optional and the defaults will be used with no arguments given.

| Flag         | Default                                | Description                                                        |
|--------------|----------------------------------------|--------------------------------------------------------------------|
| input       | "README.md"                            | The input Markdown file to be processed.                           |
| path        | "result"                                 | Indicates the output path. The default is `result`folder.          |
| title       | "Docs"                                 | Title of the documentation page.                                   |
| author      | "Project Author"                       | Author of the project to place in the 'author' meta tag.           |
| description | "Documentation of an awesome project." | Description of the project to place in the 'description' meta tag. |

An example usage with all the flags set would be as follows:
```bash
frank -input input.md -path ResultFolder -title "My Docs Title" -author "Document Author" -description "Such an amazing project"
```

To access all the flags with their descriptions, the `-h` flag can be used.

```bash
$ frank -h
Usage of ./frank:
  -author string
    	Author of the project to place in the 'author' meta tag. (default "Project Author")
  -description string
    	Description of the project to place in the 'description' meta tag. (default "Documentation of an awesome project.")
  -input string
    	The input Markdown file to be processed. (default "README.md")
  -path string
    	The path to export the resulting files. (default "result")
  -title string
    	Title of the documentation page. (default "Docs")
```

## Contributing

Frank uses dep to manage its dependencies. Even though it has a very simple dependency tree, I found it best to use a dependency management tool. Therefore, after cloning, you need to `dep ensure` to get the missing dependencies. Also, Frank uses (go-bindata)[https://github.com/jteeuwen/go-bindata] for embedding the static assets to the singular binary, which means that in order to start developing, you need to have `go-bindata` executable in your `$GOPATH/bin`.

```bash

# Clone the repository.
git clone https://github.com/karakanb/frank.git

# Install the dependencies using dep.
cd frank
dep ensure

# Generate the embedded assets for the binary file.
go generate

# Build the source.
go build

# Run the program by giving the input markdown path.
./frank -input input.md
```
