package main

import "fmt"

func main() {
	add := func(x, y int) int {
    	return x + y
	}

	val := add(10, 20)
	fmt.Println(val)

	fmt.Println("---")

	init := 0
	fmt.Println("Iniital Value", init)
	increment := func() int {
    	init++
    	return init
	}

	fmt.Println("Incremented ", increment())
	fmt.Println("Incremented ", increment())

	fmt.Println("---")

	adder := func(x int) func(int) int {
    	return func(y int) int {
        	return x + y
    	}
	}

	addFive := adder(5)
    
	// Here addFive is a fucntion, which has a set value of x=5.
	// So now, any invocation of addFive with a parameter will set the value of y
    
	// addFive will be the function
	// func (y int) int {
    // 	return 5 + y
	// }
    
    
	fmt.Println(addFive(100))
    
	addTen := adder(10)
	// Similarly, addTen is the function as follows
	// func (y int) int {
    // 	return 10 + y
	// }

	fmt.Println(addTen(100))


	fmt.Println("---")


	multiplier := func(factor int) func(int) int {
    	return func(y int) int {
        	return factor * y
    	}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println("Double 5", double(5))
	fmt.Println("Triple 10", triple(10))

}