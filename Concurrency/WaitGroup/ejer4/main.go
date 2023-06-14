package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		random := rand.Intn(150-100) + 100
		wg.Add(1)
		go func(f1 float64, wg *sync.WaitGroup) {
			x := math.Sqrt(f1)
			fmt.Println(x)
			wg.Done()
		}(float64(random), &wg)
	}
	wg.Wait()
}
