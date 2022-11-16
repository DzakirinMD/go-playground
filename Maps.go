package main

import "fmt"

func Maps() {

	/*
		This will create a map with string -> float64  (key is string format, value is float64 format).
		The last key,value will also need comma. If not it will get error
	*/
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffe pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, " - ", value)
	}

	// ints as key type
	phonebook := map[int]string{
		12394984: "maria",
		85847493: "luigo",
		84949593: "pinch",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[12394984])

	phonebook[12394984] = "bowser"
	phonebook[84949593] = "yoshi"

	for key, value := range phonebook {
		fmt.Println(key, " - ", value)
	}

}
