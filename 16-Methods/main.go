package main

import "fmt"



func main() {

   t1 :=User{"Taslim" ,"taslim@gmail.com" , true, 56}
   // { } curly braces are used here instead of normal braces
	fmt.Println(t1)
	fmt.Printf("Details are %+v", t1) //For named results
	fmt.Println(t1.Name)

	t1.getStatus()
	t1.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	
}

//Method of struct
func (u User) getStatus(){
	fmt.Println("The status is ", u.Status)
}

//When struct is passed on to a functio a new copy is passed so the defualt value cannot be changed like this
//So pass it as pointers if you wnt to change some value
func (u User) NewMail(){
	u.Email = "xyz@gmail.com"
	fmt.Println("New Email is ", u.Email)
}