package main

import "fmt"

// Global variable 'g'
var g = 1000

func main() {
	var str string = "Hello World"
	fmt.Println(str)

	str = "Happy"
	fmt.Println(str)

	fmt.Println(str + " and Excited")

	x := 100
	fmt.Println(x)

	var y = 200
	fmt.Println(x+y)

	fmt.Println(g)

	//calling another function from the main function
	newFunction()
}

func newFunction() {
	// Even this function has access to the Global Variable 'g'
	fmt.Println(g)
}
