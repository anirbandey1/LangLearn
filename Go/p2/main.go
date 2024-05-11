package main

import (
	"fmt"
)

func main() {
	slice1 := []int{}
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 3)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Printf("Length = %d\n", len(slice1))
	fmt.Printf("Capacity = %d\n", cap(slice1))
	fmt.Printf("Value = %v\n", slice1)

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Value at index %d = %d\n", i, slice1[i])
	}

	fmt.Println("For Each Loop")
	for idx, val := range slice1 {
		fmt.Printf("%d -> %d\n", idx, val)
	}

	fmt.Println("For Each loop without indices")
	fmt.Print("Values : ")
	for _, val := range slice1 {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Println("For Each loop with values")
	fmt.Print("Indices : ")
	for idx, _ := range slice1 {
		fmt.Printf("%d ", idx)
	}
	fmt.Println()

}
