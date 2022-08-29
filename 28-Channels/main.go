package main

import (
	"fmt"
	"sync"
)

func main() {
	//creating a channel
	myCh := make(chan int,2) //giving an intiger num will make more burfferd channels
    wg := &sync.WaitGroup{}
	// myCh <- 5 //Adding 5 values to the channel
	// fmt.Println(<-myCh)
    wg.Add(2)
	go func (ch chan<- int, wg *sync.WaitGroup)  {
		val, isChannelOpen := <-myCh
		fmt.Println(val)
		fmt.Println(isChannelOpen)
		wg.Done()
	}(myCh, wg)

	go func (ch chan<- int, wg *sync.WaitGroup)  {
		close(myCh)
		//myCh<-6
		//myCh<-5
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}