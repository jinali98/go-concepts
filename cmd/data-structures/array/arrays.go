package main

import "fmt"

func main() {

	// this will intialize an array with 3 elements(0, 0, 0	) with the type int32
	// as the type is int32, go will use 32 bits(4byte) to store each element. in total it allocates 12 bytes
	var intArray [3]int32
	fmt.Println(intArray)

	intArray[0] = 10
	intArray[1] = 10
	intArray[2] = 10
	fmt.Println(intArray[0])
	// to get the memory location of any element in the array, we can use the & operator
	fmt.Println(&intArray[0])
	fmt.Println(&intArray[1])
	fmt.Println(&intArray[2])

	// initalizing an array with values
	var anotherArray = [3]int32{4, 5, 6}
	anotherArray2 := [3]int32{8, 5, 6}
	fmt.Println(anotherArray)
	fmt.Println(anotherArray2)

	// SLICES are more like dynamic arrays
	// we can add or remove elements from a slice.
	// we dont need to specify the size of a slice
	var intSlice = []int32{1, 2, 3}
	// initial length of the slice is 3 and go allocates 3 slots in memory
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	// we can add elements to a slice using the append function
	intSlice = append(intSlice, 4)
	// length is now 4 and capacity is 6
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	// we can also create a slice with the make function and allocate inital capacity. 3rd argument is the capacity
	var intSlice3 = make([]int32, 1, 3)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice3), cap(intSlice3))

	intSlice3 = append(intSlice3, 2, 5, 8)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice3), cap(intSlice3))

	var peopleArray = [3]string{"John", "Jane", "Smith"}
	// iterating over a slice or an array
	for index, value := range peopleArray {
		fmt.Println(index, value)
	}
}
