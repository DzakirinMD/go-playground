package main

import (
	"embed"
	"fmt"
	"github.com/spf13/viper"
)

//This package, combined with the comment //go:embed, a compiler directive,
//tells the compiler that we intend to embed files or folders in the resulting binary.
//meaning the files will be contained in our produced binary. (go build)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//We can also embed multiple files or even folders with wildcards.
//This uses a variable of the embed.FS type, which implements a simple virtual file system.

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	println(fileString)
	println(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	println(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	println(string(content2))

	viper.AddConfigPath("C:\\workspaces\\learning\\go-playground\\go-embed\\resources")
	viper.SetConfigName("application") // Register config file name (no extension)
	viper.SetConfigType("properties")  // Look for specific type
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	port := viper.Get("prod.port")

	fmt.Println(port)
}
