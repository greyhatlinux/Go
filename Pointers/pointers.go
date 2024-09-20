package main 

import "fmt"

func make_zero(xPtr *int){
	*xPtr = 0
}

// Generally, argumets in a function is passed by making a copy of itself
// With pointers, the pointer to the variable is passed. Same concept as that of C


// Pointer is represemnted using *value_type, eg *int, *float64
// &variable -> address of the variable

// &x returns the address of x, which is *x 

func make_ten(xPtr *int) {
	*xPtr = 10
}


func main() {
	x := 100
	fmt.Println("X :", x)
	make_zero(&x)
	fmt.Println("X :", x)


	// Another way to get a pointer is 'new' 
	xPtr := new(int)

	// new(value_type) creates a continer of that value_type and return a pointer to that value

	fmt.Println("xPtr : ", *xPtr)
	make_ten(xPtr)
	fmt.Println("xPtr : ", *xPtr)

	// Go is a garbage collecting programming language, which cleans up memory autimatically when nothing is referenced to it anymore 

}