package main

import (
	"fmt"
)

func main() {
	language := make(map[string]string)

	language["js"]= "javaScript"
	language["py"]= "python"
	language["rs"]= "rust"
	language["J"]= "Java"

	fmt.Println("List of all languages", language)
	fmt.Println(language["js"])
	delete(language, "J")
	fmt.Println(language)

   // Looping of map in Go
   //like Foreach loop in JS
	for key, value := range language{
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}