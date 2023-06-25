package main

import "fmt"

func main() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	for i := 1; i <= len(nums)-2; i++ {
		fmt.Println(nums[i])
	}

	for _, v := range nums[1 : len(nums)-2] {
		fmt.Println(v)
	}
}
