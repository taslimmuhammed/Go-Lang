package main

import "fmt"

func main() {
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if  loginCount>10{
		result  ="watch out"
	} else{
	result = "Pulbic user"
   }
	fmt.Println(result)

  //initializing  + checking in the same line
  if num := 11; num>10{
	fmt.Println("Number is less than 10")
  }
}