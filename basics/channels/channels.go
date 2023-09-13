package main

import (
	"fmt"
	"sync"
)

func send(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	m := <-ch
	fmt.Println(m)
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go send(ch, &wg)
		ch <- i

	}

	wg.Wait()

}
