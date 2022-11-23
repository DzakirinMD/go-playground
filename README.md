# go-playground


```bash
# Create the package manager
$ go mod init go_playground

# maroto is require to build the marotoPDF
$ go get -u github.com/johnfercher/maroto/internal

# Run the program including all the other package
$ go run .

# Only run the main function
$ go run main.go

# package into executable
$ go build .

```