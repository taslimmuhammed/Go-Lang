package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in goLang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Value:- ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1")

	case 2:
		fmt.Println("You got 2")
	case 3:
		fmt.Println("You got 3")
		fallthrough //Acts opposit to "break" statment in C, here "break" is by default
	case 4:
		fmt.Println("You got 4")
	case 5:
		fmt.Println("You got 5")
	case 6:
		fmt.Println("You got 6")
	}
}
