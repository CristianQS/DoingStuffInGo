package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")
	}
	args := os.Args[1:]
	sum, product := 0, 1
	for _, v := range args {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		sum += num
		product *= num
	}

	fmt.Println(sum)
	fmt.Println(product)
}
