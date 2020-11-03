package main

import "fmt"

func main() {
	var slice []int
	slice3 := *new([]int)

	slice2 := append(slice, 2, 3, 4, 5)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
