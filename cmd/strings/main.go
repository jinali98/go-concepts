package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings are immutable.
	// we cannot change the value of a string after it has been created
	var charString = "Hėllo"
	var indexed = string[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range string {
		fmt.Println(i, v)
	}

	var lengthOfString = len(string)

	// this gives the length of the byte array.
	//not the length of the string in characters
	// this is because go uses utf-8 encoding
	fmt.Println(lengthOfString)

	// we can concatenate strings using the + operator
	var newString = charString + " World"
	fmt.Println(newString)

	var strSlice = [5]string{"h", "e", "l", "l", "o"}

	fmt.Println(strSlice)

	// we can use string builder to build strings and concatenate them

	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var concatString = strBuilder.String()
	fmt.Println(concatString)

}
