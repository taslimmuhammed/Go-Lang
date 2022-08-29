package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string= "https://lco.dev"

func main() {
	response, err := http.Get(url)
	handleErr(err)

	fmt.Printf("This is response %T", response)
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	handleErr(err)
	content := string(dataBytes)

	fmt.Println("The data body", content)

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
