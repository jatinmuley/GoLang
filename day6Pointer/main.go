package main

import "fmt"

func main() {
	fmt.Println("Lets study pointer")


	mypointer := 23

	var ptr = &mypointer

	fmt.Println("referece of variable is ",ptr)

	fmt.Println("Value of Pointer is ",*ptr)
	*ptr = *ptr + 2
	fmt.Println("Adding the value in pointer",*ptr)

}