package main

import (
	"fmt"
	"sync"
)

func main() {
	rounds := 10

	oddCh := make(chan struct{}, 1)
	evenCh := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		num := 1
		for i := 0; i < rounds; i++ {
			<-oddCh
			fmt.Println("from oddCh", num)
			num += 2
			evenCh <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		num := 2
		for i := 0; i < rounds; i++ {
			<-evenCh
			fmt.Println("from evenCh", num)
			num += 2
			oddCh <- struct{}{}
		}
	}()

	oddCh <- struct{}{}
	wg.Wait()
	fmt.Println("finished")
}
