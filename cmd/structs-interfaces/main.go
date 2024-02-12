package main

import "fmt"

//  structs used to create our own types

type gasEngine struct {
	mpg        uint8
	galons     uint8
	ownerInfor owner
}
type electricEngine struct {
	mpkwh      uint8
	kwh        uint8
	ownerInfor owner
}

type owner struct {
	name string
}

// structs can bind to methods. and then this methods can have access to the
// struct interface with in the function

func (e gasEngine) milesLeft() uint8 {

	var miles = e.galons * e.mpg
	return miles
}

func (e electricEngine) milesLeft() uint8 {

	var miles = e.kwh * e.mpkwh
	return miles
}

// interface
type engine interface {
	milesLeft() uint8
}

// we can pass struct methods to other functions
func canMakeIt(e engine, miles uint8) bool {

	if miles <= (e.milesLeft()) {
		fmt.Println("can make it")
		return true
	}
	fmt.Println("cannot make it")

	return false
}

func main() {
	// this will have default values for the keys as no values are there initially
	var myGasEngine gasEngine = gasEngine{
		mpg:    45,
		galons: 71,
		ownerInfor: owner{
			name: "Jane",
		},
	}
	var myElecEngine electricEngine = electricEngine{
		mpkwh: 45,
		kwh:   71,
		ownerInfor: owner{
			name: "Jane",
		},
	}

	canMakeIt(myGasEngine, 100)
	canMakeIt(myElecEngine, 10)

	fmt.Println(myGasEngine)

	fmt.Println(myGasEngine.galons, myGasEngine.mpg, myGasEngine.ownerInfor)

	var milesLeft = myGasEngine.milesLeft()

	fmt.Println(milesLeft)
}
