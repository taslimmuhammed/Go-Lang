package main

import "fmt"



func main() {

   t1 :=User{"Taslim" ,"taslim@gmail.com" , true, 56}
   // { } curly braces are used here instead of normal braces
	fmt.Println(t1)
	fmt.Printf("Details are %+v", t1) //For named results
	fmt.Println(t1.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}