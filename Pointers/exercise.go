package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func swap(x, y *int) {
	// // M1 --- with a third variable 
	// z := x
	// *x = *y
	// *y = *z

	// M2 --- without a third variable
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y

}


func main() {
	x := 100

	fmt.Println("X : ", x)
	fmt.Println("Address of X : ", &x)

	fmt.Println("---")

	namePtr := new(string)
	fmt.Println("Value of Name Pointer :", *namePtr)
	*namePtr = "Harry"
	fmt.Println("Value of Name Pointer :", *namePtr)

	fmt.Println("---")

	y := 5.3
	
	fmt.Println("y  : ", y)
	square(&y)
	fmt.Println("Squared value of y : ", y)

	fmt.Println("---")

	n1, n2 := 100, -20
	
	fmt.Println("Before Swap (n1, n2):", n1, n2)
	swap(&n1, &n2)
	fmt.Println("After Swap (n1, n2):", n1, n2)

}