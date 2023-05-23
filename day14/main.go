package main

import "fmt"

func main() {
	fmt.Println("we are in functions ")
	greeter()

	result := adder(3, 5)
	fmt.Println("result is ", result)
	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("proRes is ", proRes)
	fmt.Println("myMessage is ", myMessage)
}
func proAdder(values ...int) (int, string) {

	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro Result "
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func greeterTwo() {

	fmt.Println("Hello from greeterTwo")
}
func greeter() {
	fmt.Println("Hello from greeter")
}
