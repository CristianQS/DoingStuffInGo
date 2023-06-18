package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

func main() {
	var shoes money = 3.4567
	shoes.print()
	fmt.Print(shoes.printStr())
}
