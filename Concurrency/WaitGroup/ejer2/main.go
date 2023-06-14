package main

import (
	"fmt"
	"sync"
)

func sum(f1, f2 float64, wg *sync.WaitGroup) {
	s1 := fmt.Sprintf("%.2f", f1+f2)
	fmt.Println(s1)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sum(1.367, 1.5343, &wg)
	}
	wg.Wait()
}
