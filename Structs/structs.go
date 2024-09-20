package main

import ("fmt"; "math")


// Defining a struct 
type Circle struct {
	radius float64
}

type Rectangle struct {
	length, breadth float64
}

func circle_area(c Circle) float64 {
	return math.Pi * c.radius * c.radius
}

func rectange_area(r Rectangle) float64 {
	return r.length * r.breadth
}


func main() {

	// Initialising a struct 

	var c1 Circle
	var r1 Rectangle

	fmt.Println("Circle_1 -> r: ", c1.radius)
	fmt.Println("Area of Circle_1 : ", circle_area(c1))
	fmt.Println("Rectange_1 -> (l,b):", r1.length, r1.breadth)
	fmt.Println("Area of Rectangle_1 : ", rectange_area(r1))

	fmt.Println("---")

	// This allocated memory to all fields, sets each of them to zero value and returns a pointer to that variable
	// c2 := new(Circle)
	// r2 := new(Rectangle)

	// fmt.Println("Circle_2 -> r: ", c2.radius)
	// fmt.Println("Area of Circle_2 : ", circle_area(&c2))
	// fmt.Println("Rectange_2 -> (l,b):", r2.length, r2.breadth)
	// fmt.Println("Area of Rectangle_2 : ", rectange_area(&r2))

	fmt.Println("---")
	c3 := Circle{radius: 9.3}
	r3 := Rectangle{length: 10.2, breadth: 2.6}

	fmt.Println("Circle_3 -> r: ", c3.radius)
	fmt.Println("Area of Circle_3 : ", circle_area(c3))
	fmt.Println(
		"Rectange_3 -> (l,b):",
		r3.length,
		r3.breadth,
	)
	fmt.Println("Area of Rectangle_3 : ", rectange_area(r3))

	fmt.Println("---")


	c4 := Circle{4.5}
	r4 := Rectangle{7.4, 2.6}

	fmt.Println("Circle_4 -> r: ", c4.radius)
	fmt.Println("Area of Circle_4 : ", circle_area(c4))
	fmt.Println(
		"Rectange_4 -> (l,b):",
		r4.length,
		r4.breadth,
	)
	fmt.Println("Area of Rectangle_4 : ", rectange_area(r4))




}