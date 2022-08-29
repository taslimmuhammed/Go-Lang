package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string `json:"-"` //this field will not be displayed
	Tags     []string `json:"tags,omitempty"`//this will remove this field if there are no values
	//Be aware to avoid space between tags, and omitempty
}

func main() {
DecodeJson()
}

func EncodeJson() {
	lcoCourse := []course{
		{"React js", 999, "xyz", "xyz", []string{"web-dev", "js"}},
		{"MERN", 399, "xyz", "xyz", []string{"web-dev", "js"}},
		{"Angular", 399, "xyz", "xyz", nil},
	}
    finalJson, err := json.MarshalIndent(lcoCourse,"","\t")
	handleErr(err)
	fmt.Println(string(finalJson))


}


func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React js",
		"Price": 999,
		"Platform": "xyz",
		"tags": ["web-dev","js"]
	} `) 
//the data from the web always come in the form of bytes we here convert the json we have to bytes for simulation
	// var lcoCourse course
	// chekckValid := json.Valid(jsonDataFromWeb)
	// if chekckValid {
	// 	fmt.Println("json was valid")
	// 	json.Unmarshal(jsonDataFromWeb, &lcoCourse)
	// 	fmt.Printf("%#v\n",lcoCourse)
	// } else{
	// 	fmt.Println("json was not valid")
	// }

	//some cases where you just want to add key value 

	var myOnlineData map[string] interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
    fmt.Printf("%#v\n",myOnlineData)

	for k,v := range myOnlineData{
		fmt.Printf("%v :%v ,%T\n",k,v,v)
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}