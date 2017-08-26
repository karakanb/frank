# Frank

Frank is a simple static site generator based on Markdown, with basic documentation needs in mind. It basically takes the markdown input, renders it and creates a sample output folder, which can be deployed to a very basic static site hosting service, such as GitHub Pages.

## Usage

Currently, you need to clone the library and build it. In order to build it, you must have Go and Dep installed as well.
```bash

# Clone the repository.
git clone https://github.com/karakanb/frank.git

# Install the dependencies using dep.
cd frank
dep ensure

# Build the source.
go build

# Run the program by giving the markdown path.
./frank --input input.md
```