package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Lets play a dice game ")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)+ 1
	fmt.Println("Value of Dice ",diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("dice value is one and you can open ")
	
	case 2: fmt.Println("you can move two spot")
	case 3: fmt.Println("you can move 3 spot")
	case 4: fmt.Println("you can move 4 spot")
	fallthrough
	case 5: fmt.Println("you can move 5 spot")
	case 6: fmt.Println("you can move 6 spot and roll dice again ")
default: fmt.Println("what was that ")

	}


}