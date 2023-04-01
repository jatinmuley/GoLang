package main

import "fmt"

func main() {
	fmt.Println("If Else ")

	loginCount := 10

	var result string

	if loginCount < 10{
		result = "regular user "
	}else if loginCount >10 {
		result = " watch out "

	}else {
		result = " exactly 10 login count  "
	}

	fmt.Println(result)
	if 9%2 ==2 {
		fmt.Println("number is even")
	}else {
		fmt.Println("number is odd")
	}


	if num := 3 ; num < 10 {
		fmt.Println("number is less than10 ")
	}else{
		fmt.Println("number is not  less than 10 ")
	}

	
}