package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to  time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2022, time.August, 12, 23,23, 59,99,time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}