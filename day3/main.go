package main

import "fmt"

func main() {
	welcome := "Welcome to user input "
	fmt.Println(welcome)
	fmt.Println("Please enter your name: ")

	input, _ := reader.ReadString('\n')
	fmt.Printf("Type of the inmut is %T", input)

}