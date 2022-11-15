package main

import "fmt"

func StringFormatting() {

	age := "20"
	name := "Belvin"

	fmt.Print("Hello, \n world")

	fmt.Println("Hello !")
	fmt.Println("World !")

	/* Printf (formatted Strings) %_ = format specifier
	    q = quote
		T = print the type of the variable
	*/
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("Age variable type is : %T and name type is : %T \n", age, name)
	fmt.Printf("the score is : %f \n", 225.55)
	fmt.Printf("the score is : %0.1f \n", 225.55)

	/* Sprintf (save formatted strings)
	   Documentation : https://golangdocs.com/string-formatting-in-golang
	*/
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved String is: ", str)

}
