package main

import "fmt"

func main() {
	var (
		x = 100
		y = "Hello world"
	)

	// This is not defined, since = olny works with var(), otherwise use := 
	// z = 34.05

	z := 20.40

	fmt.Println(x, y, z)
}