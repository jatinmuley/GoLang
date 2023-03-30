package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Playing with time ")
	currenttime := time.Now()
	fmt.Println("Current time is ", currenttime)
	fmt.Println(currenttime.Format("2006/01/02 15:04:05  Monday"))
}