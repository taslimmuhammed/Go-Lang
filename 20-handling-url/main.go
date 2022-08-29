package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?cousename=reactjs&paymentid=ghtnmkzcn23124"

func main() {

	fmt.Println(myUrl)
	results,_ := url.Parse(myUrl)
	// fmt.Println(results.Scheme)
	// fmt.Println(results.Host)
	// fmt.Println(results.Path)
	// fmt.Println(results.Port())
	// fmt.Println(results.RawQuery)

	qparams := results.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _,val :=range qparams{
		fmt.Println("Param is:", val)
	}
    
	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}
   
	newUrl := partOfUrl.String()

	fmt.Println(newUrl)

}