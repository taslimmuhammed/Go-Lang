package main

import (
	"bufio"
	"fmt"
	"os"
)

const LoginKey string = "hi123455" //public
func main() {
	welcome := "Welcome user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
    
	//Comma ok OR error Ok
    fmt.Println("Hey , please enter rating for our pizza")
	input , _ := reader.ReadString('\n') // Add _ if you want to ignore the try or catch [1st or 2nd variable] 
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of input - %T", input)
} 

// //no vars sign
// numberOfUsers := 3000