package main

import "fmt"
import "math"
import "math/bits"

func main() {

	x := [6]string{"a","b","c","d","e","f"}
	fmt.Println(x[2:5])

	fmt.Println("---")

	x2 := []int {
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	smallest := math.MaxInt
	for i:=1; i<len(x2); i++ {
		if x2[i] < smallest {
			smallest = x2[i]
		}
	}
	fmt.Println(smallest)

}