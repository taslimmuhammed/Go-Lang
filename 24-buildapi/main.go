package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int  `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db
var courses []Course

//middleware, helper -file

func (c Course) isEmpty() bool {
	return  c.CourseName == ""
}
func main() {
	greeter()
	r := mux.NewRouter()

	//seeding --means adding some values
	courses = append(courses, Course{CourseId: "1",CourseName: "React Js", CoursePrice: 299, Author: &Author{Fullname: "name_author", Website: "go.dev"}})
 
	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/cources",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCource).Methods("GET") //this {id} is called using the params 
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOnCource).Methods("PUT")
	r.HandleFunc("/course/{id}",deletOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000",r))

  }
  
  func greeter() {
	  fmt.Println("hello there mod users")
  }

func serveHome(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("<h1>Hello there</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Request to get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCource(w http.ResponseWriter, r *http.Request){
	fmt.Println("get one ocurse")
	w.Header().Set("Content-Type", "application/json") //setting headers for api
	
	//grab Id from request [from url]
	params:= mux.Vars(r)

	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
     str := "hello" +params["id"]
	 fmt.Println(str)
	json.NewEncoder(w).Encode(str)
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("get one ocurse")
	w.Header().Set("Content-Type", "application/json")

	//what if : body is empty
	if r.Body==nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course

	json.NewDecoder(r.Body).Decode(&course) //Decode reads the next JSON-encoded value from its input and stores it in the value pointed . 

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside")
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) //convert it  to string
	courses = append(courses, course)  
	json.NewEncoder(w).Encode(course)
	return
}

func updateOnCource(w http.ResponseWriter, r *http.Request){
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")

    // grab id from req
	params:= mux.Vars(r)

	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)// to remove the exact element from array
			_ = json.NewDecoder(r.Body).Decode(&course)
           course.CourseId = params["id"]
		   courses = append(courses, course)
		   json.NewEncoder(w).Encode(course)
		   return
		}
	}
}

func deletOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

    for index, course := range courses{
		if course.CourseId == params["id"]{
             courses = append(courses[:index],courses[index+1:]...)
			 break;
		}
	}
}