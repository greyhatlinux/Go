package main

import "fmt"

func main() {
	fmt.Println("In this we'll learn about loops and other control structures in Go")

	// i :=1 
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// fmt.Println(" ---- ")
	
	for i:=1; i<=10; i++ {
		if (i%2 == 0){
			fmt.Println(i, " -> Even")
		} else {
			fmt.Println(i, " -> Odd")
		}
	}

}
