package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f", m)
}

func main() {
	var shoes money = 3.4567
	shoes.print()
}
