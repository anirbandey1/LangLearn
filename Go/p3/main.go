package main

import (
	"fmt"
)

func main() {
	fmt.Println("2-d arrays")

	matrix1 := [3][3]int{}

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[i]); j++ {
			fmt.Printf("%d ", matrix1[i][j])
		}
		fmt.Println()
	}

	var nrows2 int = 4
	var ncols2 int = 4

	matrix2 := make([][]int, nrows2)
	for i := range matrix2 {
		matrix2[i] = make([]int, ncols2)
	}

	fmt.Printf("Initial Value of Matrix 2 = %v\n", matrix2)

	for i := range matrix2 {
		for j := range matrix2[i] {
			matrix2[i][j] = i + j
		}
	}

	fmt.Printf("Updated first time - Value of Matrix 2 = %v\n", matrix2)
}
