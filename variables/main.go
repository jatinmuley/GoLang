package main

import "fmt"

// When first letter of the variable is Caps its an Public variable
const LoginToken string = "Jfffffff"
func main() {
	var username string = "Jatin"
	fmt.Printf(username)
	fmt.Printf(" Variable is of type  : %T \n",username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf(" Variable is of type  : %T \n",isLoggedIn)

	// implicit type 
	// lexer is deciding the type

	var anything = " jatin Muley "
	fmt.Println(anything)
// without the var only allowed inside a method 
	numberofuser := 300000
	fmt.Println(numberofuser)


	fmt.Println(LoginToken)
	fmt.Printf(" Variable is of type  : %T \n",LoginToken)

}