package main

import "fmt"

func main() {
	fmt.Println("Structs in GO Lang ")
	// there is no inheritance 
	// no class 

jatin := User{"Jatin","Jatin@gmail" , true ,23}
fmt.Println(jatin)

fmt.Printf("Details of JAtin are %+v\n",jatin)
fmt.Printf("Name is %v and Email is %v",jatin.Name,jatin.Email)
}


type User struct {

	Name string 
	Email string 
	status bool
	Age int 
}