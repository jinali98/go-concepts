package main

import "fmt"

func main() {

	// by default this will be int32 or int64 depending on the system
	var num1 int = 5;
	var num2 int8 = 10;
	// var num3 int16 = 15;

	// // unsigend int stores only positive values
	// var num4 uint = 20;

	// // float32 and float64. 
	// var num5 float32 = 3.14;

	// //we should use float64 as it is more precise.
	// var num6 float64 = 3.14;


	// we cannot do math operations on different types
	// we need to convert them to the same type(casting)
	// var result = num1 + num2; // this will give an error
	var resultOne = num1 + int(num2);


	fmt.Println(resultOne)

	
}