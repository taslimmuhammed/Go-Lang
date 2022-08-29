package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
  greeter()
  r := mux.NewRouter()
  r.HandleFunc("/", serveHome).Methods("Get")

//--second Way
log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hello there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request){
   w.Write([]byte("<h1>Hello from new Server</h1>"))
}