package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(f1 float64, wg *sync.WaitGroup) {
		x := math.Sqrt(f1)
		fmt.Println(x)
		wg.Done()
	}(1.34, &wg)
	wg.Wait()
}
