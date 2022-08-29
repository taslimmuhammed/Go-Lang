package main

import (
	"fmt"
	"sync"
)

func main() {
	
	//func (){}() //anonymus functions inside Go lang
    
	wg := &sync.WaitGroup{}
    mut := &sync.RWMutex{}
    
	mut.RLock()
	var score = []int{0}
    mut.RUnlock()

	wg.Add(3) //Letting Go know to wait untill 3 threads return 
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	mut.RLock()
	fmt.Println(score)
	mut.RUnlock()
}
