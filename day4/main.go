package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our app");
fmt.Println("Give us Rating ")
	reader := bufio.NewReader(os.Stdin)
	input ,  _  := reader.ReadString('\n')
	fmt.Println("your rating  - ",input )

	numrating   , err  :=  strconv.ParseFloat(strings.TrimSpace(input),64)
	if (err != nil){

		fmt.Println(err)
	}else{
		fmt.Println("Added to the rating ", + numrating + 1)
	}
}