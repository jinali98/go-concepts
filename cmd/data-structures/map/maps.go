package main

import "fmt"

func main() {

	var map1 = make(map[string]uint8)
	map1["one"] = 1
	map1["two"] = 2

	var map2 = map[string]uint8{
		"one": 1,
	}

	// map returns 2 values. 1st is the value and 2nd is a boolean indicating if the key was found
	var number, ok = map2["two"]

	if ok {
		fmt.Println(number)
	} else {
		println("key not found")
	}

	var map3 = map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Println(map3[1])

	// deleting a key from a map.
	// this is not gonna return anything
	delete(map3, 1)

	var peopleMap = map[string]uint8{
		"John":  25,
		"Jane":  30,
		"Smith": 35,
	}

	// iterating over a map
	// elements do not have an order in a map.
	//so the order of the elements each time we run that in a loop is not the same

	for key := range peopleMap {
		fmt.Println(key)
	}

	for key, value := range peopleMap {
		fmt.Println(key, value)
	}
}
