# go-playground

This playground consist of basic go understanding. 
There are 2 mini project to make further understand in to the concept of...

1. Struct, Pointer and User Input -> bills
2. Goroutines, Channel, Buffered Channel and Waitgroups -> goroutine
3. FileIO, PDF Generation 


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

<h1>How to run goconcurrency project</h1>

```bash
# go into goconcurrency folder
$ cd goconcurrency

# run the program

$ go run .

```