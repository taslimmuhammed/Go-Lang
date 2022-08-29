package main

import "fmt"

func main() {
	myNumber := 23
	var ptr   = &myNumber

	fmt.Println(ptr, *ptr)

}