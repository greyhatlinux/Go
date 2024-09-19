package main

import "fmt"

func main() {
	// Maps -> unordered collection of key:value pairs
	// Used to look up the value(s) associated with the key 

	// Syntax - map[key_type]value_type

	mp := make(map[int]string)

	mp[1] = "apple"
	mp[2] = "banana"
	mp[4] = "mango"
	mp[10] = "guava"

	fmt.Println(mp)
	
	// for unavailable key-value pairs
	fmt.Println("map[100]->", mp[100])

	// deleting from maps
	delete(mp, 10)

	fmt.Println(mp)

	// Accessing map elements can return 2 values, 
	// 1st value -> result of the lookup
	// 2nd value -> whether the lookup was successful or not (1,0)

	// fruit, exists := mp[12]
	// fmt.Println(fruit, exists)

	if fruit, exists := mp[2]; exists {
		fmt.Println(fruit, exists)
	}

	if fruit2, exists := mp[200]; exists {
		fmt.Println(fruit2, exists)
	}

	fmt.Println("---")

	//  Declaring a map, using the 'make' keyword
	mp2 := make(map[int]string)

	// Using the map, after it is declared 
	mp2[2] = "Alpha"

	fmt.Println(mp2)

	// Declaring and using at the same point doesn't require make keyword
	mp3 := map[int]string {
		2 : "Phone",
		3 : "Laptop",
		4 : "Earbuds",
	}

	fmt.Println(mp3)

	fmt.Println("---")

	// Nested Map

	elements := map[string]map[string]string {
		"H" : map[string]string{
			"name": "Hydrogen",
			"state": "gas",
		},
		"He" : map[string]string {
			"name" : "Helium",
			"state": "gas",
		},
		"Li": map[string]string {
			"name": "Lithium",
			"state": "solid",
		},
	}

	fmt.Println(elements)

	if el, exists := elements["Li"]; exists {
		fmt.Println(el)
	}

	if el2, exists := elements["C"]; exists {
		fmt.Println(el2)
	} else {
		fmt.Println("C doesn't exists")
	}


}