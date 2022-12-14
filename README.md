# go-playground

This playground consist of basic go understanding. 
There are 2 mini project to make further understand in to the concept of...

1. <b>bills</b> -> Struct, Pointer and User Input
2. <b>goconcurrency</b> -> Goroutines, Channel, Buffered Channel and Waitgroups
3. <b>marotoPDF</b> -> FileIO, PDF Generation 
4. <b>Port Scanner</b> -> Goroutine, Mutex, Waitgroup
5. <b>ATM</b> -> Mutex, Goroutine, Waitgroup
6. <b>Quiz App</b> -> Timer, File reader
7. <b>Weather tracker</b> -> For consuming REST API


```bash
# Create the package manager
$ go mod init go_playground

# update go package
$ go mod tidy


# Run the program including all the other package
$ go run .

# Only run the main function
$ go run main.go

# package into executable
$ go build .

```
<h2>How to run bills project</h2>

from this project you will get to use features such as Struct, Pointer and User Input

```bash
# run the program
$ go run main.go
```


<h2>How to run goconcurrency project</h2>

from this project you will get to use features such as Goroutines, Channel, Buffered Channel and Waitgroups

```bash
# go into goconcurrency folder
$ cd goconcurrency

# run the program

$ go run main.go
```

<h2>How to run marotoPDF project</h2>

from this project you will get to use features such as FileIO, PDF Generation

```bash
# go into marotoPDF folder
$ cd mini-project/marotoPDF

# run the program

$ go run main.go
```

<h2>How to run Port Scanner project</h2>

from this project you will get to use features such as FileIO, PDF Generation

```bash
# go into port-scanning folder
$ cd mini-project/port-scanning

# run the program

$ go run main.go
```

<h2>How to run ATM project</h2>

from this project you will get to use features such as Mutex, Goroutine, Waitgroup

```bash
# go into port-scanning folder
$ cd mini-project/port-scanning

# run the program

$ go run main.go
```

<h2>How to run weather tracker project</h2>

from this project you will get to use consuming REST API feature

```bash
# go into port-scanning folder
$ cd mini-project/weather-tracker

# run the program

$ go run main.go
```

to test the api open browser and use below url:
1. [test hello go](http://localhost:8080/hello)
2. [singapore temp](http://localhost:8080/weather/singapore) or eg: http://localhost:8080/weather/*city*