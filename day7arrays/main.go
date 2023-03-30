package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go lang ")

	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[3] = "Orange"

	fmt.Println("Fruit list is ",fruitlist)
}