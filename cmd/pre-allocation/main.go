package main

import (
	"fmt"
	"time"
)

// pre allocation of capacity to a slice can improve performance

func main() {
	var n = 1000000
	var slice = []int{}
	// pre allocate capacity to the slice
	var slice2 = make([]int, 0, n)

	fmt.Println("Length of slice 1: ", len(slice))
	fmt.Println("Capacity of slice 1: ", cap(slice))
	fmt.Println("Length of slice 2: ", len(slice2))
	fmt.Println("Capacity of slice 2: ", cap(slice2))

	fmt.Println("Time taken to append to slice 1 without pre capacity allocation: ", timeLoop(slice, n))
	fmt.Println("Time taken to append to slice 2 with pre capacity allocation: ", timeLoop(slice2, n))

}

func timeLoop(slice []int, n int) time.Duration {

	var t0 = time.Now()

	for len((slice)) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}
