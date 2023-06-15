package main

import "fmt"

func main() {
	var c1 chan float64
	c2 := make(<-chan rune)
	c3 := make(chan<- rune)
	c4 := make(chan int, 10)
	fmt.Printf("%T, %T, %T, %T\n", c1, c2, c3, c4)
}
