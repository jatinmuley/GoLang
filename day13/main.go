package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in Golang")

	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday"}
	fmt.Println(days)
	// for i := 0; i < len(days) ; i++{
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index , day := range days{
	// 	fmt.Printf("index is  %v and value is %v",index , day)
	// }
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {

			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue

		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("I am at lco")
}
