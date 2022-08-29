package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	//"math/rand"
	//"time"
)

func main() {
	// var num1 int = 2
	// var num2 float64 = 5

	//random number through math/random package
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(50))
	

	//random from crypto package
	for i:=0; i<20;i++{
		myRandNum ,_ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Println(myRandNum)
	}
	}
	