# go-playground

This playground consist of basic go understanding. 
There are 2 mini project to make further understand in to the concept of...

1. Struct, Pointer and User Input -> bills
2. Goroutines, Channel, Buffered Channel and Waitgroups -> goconcurrency
3. FileIO, PDF Generation -> marotoPDF


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
<h1>How to run bills project</h1>

```bash
# go into bills folder
$ cd bills

# run the program

$ go run .
```


<h1>How to run goconcurrency project</h1>

```bash
# go into goconcurrency folder
$ cd goconcurrency

# run the program

$ go run .
```

<h1>How to run marotoPDF project</h1>

```bash
# go into marotoPDF folder
$ cd marotoPDF

# run the program

$ go run .
```