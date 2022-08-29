package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
    // GetRequest()
	//JSONPostRequest()
	PostFormRequest()
}


func PostFormRequest() {
	const myUrl = "http://localhost:8080/postform"

	//formdata

	data := url.Values{}

	data.Add("firstname","Taslim")
	data.Add("age","21")
    
	response, err := http.PostForm(myUrl, data)
	handleErr(err)
    defer response.Body.Close()
    content, err := ioutil.ReadAll(response.Body)
    handleErr(err)
    fmt.Println("Data", string(content))
}

func JSONPostRequest() {
   const myUrl = "http://localhost:8080/post"
   //strings.NewReader is used to create any type of data
   //using here to create a fake json data
   requestBody := strings.NewReader(`{
	"coursename":"GO lang ",
	"age":"21"
   }`)

   response, err := http.Post(myUrl, "application/json", requestBody)
   handleErr(err)
   content, err := ioutil.ReadAll(response.Body)
   handleErr(err)
   fmt.Println("Data", string(content))
}



func GetRequest() {
	const myUrl = "http://localhost:8080"
	response, err := http.Get(myUrl)

	handleErr(err)
	defer response.Body.Close()

	fmt.Println("Status: ",response.Status)
	fmt.Println("Cotent length:" , response.ContentLength)
    
    dataBytes,err := ioutil.ReadAll(response.Body)

	//Here are two methods to print the incoming data
	// fmt.Println("Data : " , string(dataBytes))

	var responseString strings.Builder
	byteCount,_ := responseString.Write(dataBytes)
	fmt.Println(byteCount)
	fmt.Println("Data : " ,responseString.String())

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}