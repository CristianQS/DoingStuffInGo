package main

import "fmt"

func main() {
	fruits := []string{"banana", "apple", "orange"}
	for i, value := range fruits {
		fmt.Printf("Index: %d, Element: %q\n", i, value)
	}

}
