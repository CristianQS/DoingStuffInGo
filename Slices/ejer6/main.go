package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	friendsCopy := make([]string, len(friends))
	copy(friendsCopy, friends)

	friends[0] = "Pepe"

	fmt.Printf("%v", friends)
	fmt.Printf("%v", friendsCopy)

}
