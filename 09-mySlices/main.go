package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple","Tomato","Peach"}
	fmt.Printf("Type Of fruit list is %T \n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3]) //start from 1 till 3
	fmt.Println(fruitList)

	highScroes := make([]int, 4)
    highScroes[0] = 23
	highScroes[1] = 89
	highScroes[2] = 70
	highScroes[3] = 63

	//scores := []int{2, 3, 5}
    
	highScroes = append(highScroes, 53, 79) //When append is called new array is crated and reAllocated
	fmt.Println(highScroes)
    sort.Ints(highScroes)
	fmt.Println(highScroes)
	highScroes = append(highScroes[:2], highScroes[3:]...,) // to skip the values   ... - is to add more value to append
	//cant concatinates two or more slices at the same time
	fmt.Println(highScroes)
}