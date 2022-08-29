package main


import "fmt"

type cube interface {
	//methods
	Area() int
	Voleum() int
}
//you can only declare methods iside an interface
type myValue struct {
	length int
	width  int
	hieght int
}

func (m myValue) Area() int {
	return m.length * m.width
}

func (m myValue) Voleum() int {
	return m.hieght * m.length * m.width
}

func main() {
	var c1 cube
	c1 = myValue{1, 2, 3}
	fmt.Println("The area of the cube is", c1.Area())
	fmt.Println("The voleum of the cube is", c1.Voleum())

}