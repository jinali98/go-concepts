package main

func main() {

	var intArray = [5]int32{1, 2, 3, 4, 5}

	// itering over an array using for loop

	for i, v := range intArray {
		println(i, v)
	}

	// golang doesnt have while or do-while loops.
	// we can achive the same using for loop as below

	var limit = 0

	for limit < 10 {
		println(limit)
		limit++
	}

	for {

		if limit < 10 {
			break
		}

		println(limit)
		limit++
	}

	// another for loop approach
	for i := 0; i < 10; i++ {
		println(i)
	}
}
