package main

import "fmt"

func main() {
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
