package main

import "fmt"

func first_func() {
	fmt.Println("First Function")
}

func second_func() {
	fmt.Println("Second Function")
}

func third_func() {
	fmt.Println("Third Function")
}

func fourth_func(){
	fmt.Println("Fourth Function")
}


func main() {

	// first_func()

	// // Defering a function will make it run at the end of the program.
	// defer second_func()

	// third_func()
	// fourth_func()


	// // This function ultimately becomes :
	// // first_func()
	// // third_func()
	// // second_func() -> defered function at the end

	fmt.Println("---")

	first_func()
	second_func()
	defer third_func()
	defer fourth_func()

	// This program becomes :
	// first_func()
	// second_func()
	// fourth_func()
	// third_func()

	// Defered functions are called bottom to top

}