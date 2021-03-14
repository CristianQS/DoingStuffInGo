package main

import (
	"fmt"
	"strconv"
)

func main()  {
	for number := 1; number < 100; number++ {
		var result string
		if isMultiple3(number) {
			result += "Fizz"
		}
		if isMultiple5(number) {
			result += "Buzz"
		}
		if result == "" { 
			result = strconv.Itoa(number)	
		}
		fmt.Println(result)
	}
	
}

func isMultiple3(number int)(bool) {
	return number % 3 == 0
}

func isMultiple5(number int)(bool) {
	return number % 5 == 0
}