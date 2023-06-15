package main

import "fmt"

func main() {
	c := make(chan string)
	go func(s string) {
		c <- s
	}("Hello")
	value := <-c
	fmt.Println("Value received from channel:", value)
}
