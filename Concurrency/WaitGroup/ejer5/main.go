package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	*b += n
	mutex.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	*b -= n
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &mutex)
		go withdraw(&balance, i, &wg, &mutex)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
