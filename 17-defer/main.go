package main

import "fmt"

func main() {
	defer fmt.Println("hi ")
	defer fmt.Println("hello")
	 fmt.Println("how")
	 fmt.Println("are you")
	 myDiffer()
}

func myDiffer(){
	for i :=0 ; i<5; i++ {
		defer fmt.Println(i)
	}
}