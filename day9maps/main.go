package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLAng ")
	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["Py"] = "python"
	languages["C"] = "C"

	fmt.Println("Values of Maps is ",languages)
	fmt.Println("JS is for ",languages["JS"])

	// delete value from maps 
	delete(languages, "JS")

	fmt.Println("The languages are ",languages)

	// loops 

	for key , value := range languages {
		fmt.Println("Key is ",key," and Value is ",value)
		
	}
}