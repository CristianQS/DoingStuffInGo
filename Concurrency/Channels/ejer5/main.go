package main

import "fmt"

func power(n int, c chan int) {
	c <- n * n
}

func main() {
	c := make(chan int)
	for i := 0; i < 50; i++ {
		go func(n int) {
			c <- n * n
		}(i)
		fmt.Println(<-c)
	}
}
