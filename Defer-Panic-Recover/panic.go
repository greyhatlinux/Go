package main

import "fmt"

// Panic and recover is error handling mechanism in Go


func main() {
	// Panic generally indicated programmer error,
	// eg, accessing out of bound index of an array

	for i:=0;i<5;i++ {
		fmt.Println(i+1)
	}
	
	fmt.Println("Hello World!")


	// Panic can be paired with defer 
	// panic("Error!!")

	defer func() {
		str := recover()
		fmt.Println(str)
	}() 
	panic("Panic!!")


}