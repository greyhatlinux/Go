package main

import "fmt"

func main() {
	// Like arraya, slices are indexable and lave length
	// Unlike arrays, slices are allowed to change

	var arr []float64
	// Notice the missing length between the []. This makes it slices, and not array
	// Default size of arr is 0
	fmt.Println("Arr1	:", arr)

	arr2 := make([]float64, 5)
	fmt.Println("Arr2	:", arr2)

	fmt.Println("---")

	arr3 := make([]float64,  5, 15)
	// Original size of the array = 5
	// But the actual size of the array created by the system is 15. So It won't create a new array uptill capacity of 15
	
	arr3[3] = 300
	fmt.Println("Arr3			: ",arr3)

	fmt.Println("Length of Arr3		:",len(arr3))
	fmt.Println("Capacity of Arr3	:",cap(arr3))

	arr3 = append(arr3, 600, 700)

	fmt.Println("Arr3			:",arr3)	

	// We cannot go beyond the lenth, to alter elments. It has to be within the length
	// arr3[12] = 1200
	// fmt.Println("Arr3 with 12th element	:", arr3)

	fmt.Println("---")

	big_arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	small_arr := big_arr[2:5]

	fmt.Println("Original array	", big_arr)
	fmt.Println("Sliced array	", small_arr)

	fmt.Println("---")

	// Append and copy functions
	arr4 := []int {1,2,3,4}
	arr5 := append(arr4, 500,600)

	fmt.Println("Original Array ", arr4)
	fmt.Println("Appended Array ", arr5)

	fmt.Println("---")

	arr6 := []int{1,2,3,4}
	arr7 := make([]int, 3)

	copy(arr7, arr6)

	// copy (dest, src) copies from source to destination. 
	// Since arr7 can have only 3 elements, the first 3 elements are copied from arr6 to arr7

	fmt.Println("Original Array ", arr6)
	fmt.Println("Copied Array   ", arr7)

}
