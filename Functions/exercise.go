package main

import "fmt"

func add(nums []int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func even_or_odd(x int) (half int, is_even bool){
	if x%2 == 0 {
		return x/2, true
	}
	return x/2, false
}

func greatest_function(x ...int) int {
	max := -1
	for _, value := range x{
		if value > max {
			max = value
		}
	}
	return max
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)

}

func main() {

	arr := []int {1,2,3,4,5,6,7,8,9}
	fmt.Println(arr)
	fmt.Println("Add arr[1,5] : ", add(arr[1:5]))
	fmt.Println("Add arr[4:] : ", add(arr[4:]))

	fmt.Println("---")


	v10,b10 := even_or_odd(10)
	fmt.Println(v10, b10)

	fmt.Println(even_or_odd(17))

	fmt.Println("---")

	arr2 := []int{1,2,3,4,56,7,44,99,0}
	fmt.Println(arr2)
	fmt.Println(greatest_function(1,2,3,4,56,7,44,99,0,))
	
	fmt.Println("---")

	fmt.Println("Fibonacci 10 ", fibonacci(10))
	fmt.Println("Fibonacci 15 ", fibonacci(15))


}