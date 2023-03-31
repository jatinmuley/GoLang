package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices ")
	var fruitlist = []string{"mango","banana","apple"}
	fmt.Println("fruitlist is ",fruitlist)
fruitlist = append(fruitlist, "peach","kiwi")
fmt.Println("fruitlist is ",fruitlist)

// slicing the slice 
fruitlist = append(fruitlist[1:])
fmt.Println(fruitlist)

fruitlist = append(fruitlist[1:3])
fmt.Println(fruitlist)


highscore := make([]int,4 )
highscore[0] = 232
highscore[1] = 134
highscore[2] = 534
highscore[3] = 634

fmt.Println("highscore is ",highscore)
highscore = append(highscore , 232,333,3334)
fmt.Println("Highsocre updates",highscore)
sort.Ints(highscore)
fmt.Println(highscore)
 fmt.Println(sort.IntsAreSorted(highscore))


// remove a value based on index
 var cources = []string{"react","js","Python","ruby","Java"}
 fmt.Println(cources)
 var index =2 
 cources = append(cources[:index],cources[index+1:]...)
 fmt.Println(cources)
}