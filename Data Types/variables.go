package main

import "fmt"

func main() {

	fmt.Println("1 + 5 = ", 1 + 5)


	fmt.Println("1.0 + 5.0 = " , 1.0 + 5.0)
	fmt.Println("1.1 + 5.0 = " , 1.1 + 5.0)


	// The 2nd character of the word "Hello World' -> e in bytes is 101
	fmt.Println("Hello World"[1])

	fmt.Println("Hello " + "World")


	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!(false && true && true))
	fmt.Println(!true)
	fmt.Println(!!!false)

}