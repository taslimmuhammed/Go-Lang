package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Thursday"}
	fmt.Println(days)

	//For Loops
	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

//Using range keyWord
	// for i:= range days{ //The rage automatically lopps through array or slices And I is the Index
    //    fmt.Println(days[i])
	// }

	for index, day :=range days{
		fmt.Printf("index is %v and value is %v \n", index, day)
	}
	//An "_" can be used to ingore either of these values

    var rougeValue int = 1
    for rougeValue <10 {
		rougeValue++
		if rougeValue ==3 {goto lco}
		if rougeValue ==5 {break} //break or countinue is vald here
		fmt.Println(rougeValue)
	}

	lco : 
	fmt.Println("Skippe dto these part")

	


}