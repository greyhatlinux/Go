package main

import "fmt"

func main() {
	var arr[5] int 
	arr[4] = 100
	arr[2] = 20
	fmt.Println(arr)

	total := 0

	for i:=1;i<5;i++{
		total += arr[i]
	}
	fmt.Println("Total sum = ", total)
	fmt.Println("Average   = ", total/len(arr))

	fmt.Println("---")

	var arr2[5] float32
	arr2[1] = 163
	arr2[2] = 250
	arr2[4] = 443
	arr2[3] = 371

	fmt.Println(arr2)

	var total2 float32 = 0

	for _, value := range arr2 {
		total2 += value
	}

	fmt.Println("Total 	 = ", total2)
	fmt.Println("Average = ", total2/float32(len(arr2)))

	fmt.Println("---")

	// Another way to declare array in Go
	arr3 := [5]float64{10,20,30,40,50,}

	fmt.Println(arr3)

}