package main

import "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	mySlice = append(mySlice, 6.0)

	a := 10.0
	mySlice[0] = a

	mySlice = append(mySlice, 10.10)

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)

}
