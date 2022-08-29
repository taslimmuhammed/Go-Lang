//initiation
//go mod init <name>

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointers
var signals = []string{"test"}
var mut sync.Mutex  //mutex lock
func main(){
//    go greeter("helloo") //Go keyword fires-up a new thread
//    greeter("world")

	websiteList := []string{
     "http://lco.dev",
	 "http://go.dev",
	 "http://google.com",
	 "http://fb.com",
	 "http://github.com",
	}

	for _, site := range websiteList{
       go getStatusCode(site)
	   wg.Add(1) //adds total number of threads 
	}

	wg.Wait() //waits for all threads to end 
    fmt.Println(signals)
}

// func greeter(s string){
// 	for i:=0;i<5;i++{
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil{
		fmt.Println("Error in endpoint -", err)
	}else{
		fmt.Printf("%d status code for %s \n", res.StatusCode,endpoint)
		mut.Lock()
		signals = append(signals,endpoint)
    	mut.Unlock()	
	}
	
}

//3 function of waitGroup -wg
//wg.Add(1), wg.Wait(), wg.Done()