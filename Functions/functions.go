package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	n := []float64{1, 24, 3, 41, 56, 64}
	fmt.Println(avg(n))

	n1 := []float64{1, 21, 33, 42, 5, 63, 7}
	total := 0.0

	for _, v := range n1 {
    	total += v
	}

	fmt.Println(total / float64(len(n1)))

	fmt.Println("Sum Total : ", sum(n1))

	//fmt.Println("Max and min values ", max_min_values(n1))
	mx, mn := max_min_values(n1)
	fmt.Println(mx, mn)

}

func avg(nums []float64) float64 {
	var sum float64
	for _, value := range nums {
    	sum += value
	}
	return sum / float64(len(nums))
}

func sum(nums []float64) (sum_total float64) {
	sum_total = 0.0
	for _, num := range nums {
    	sum_total += num
	}
	return
}

func max_min_values(nums []float64) (float64, float64) {
	max := 0.0
	min := 999999.0

	for _, value := range nums {
    	if value > max {
        	max = value
    	}
    	if value < min {
        	min = value
    	}
	}

	return max, min
}

