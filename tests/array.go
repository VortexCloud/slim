package main

import "fmt"

func main() {

	arr := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	var multi [9][9]string

	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 {
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
		}
	}

	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2)
		}
		fmt.Println()
	}
}
