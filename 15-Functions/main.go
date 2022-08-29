package main

import "fmt"

func main() {
	result := adder(1, 2)
	fmt.Println(result)
	res2, _ := proAdder(1, 2, 4, 5)
	fmt.Println(res2)
}

func adder(a int, b int) int{
   return a+b 
}

func proAdder(values ...int) (int, string){
	total := 0
	for _, val := range values{
		total +=val
	}
	return total , "this is a message"
}