package main

import "fmt"

//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

type Rectangle struct {
	Length int
	Width  int
}

func area(r Rectangle) int {
	return r.Length * r.Width
}

func perimeter(r Rectangle) int {
	return 2 * (r.Length + r.Width)
}

func main() {
	r := Rectangle{Length: 10, Width: 5}
	fmt.Println("Area: ", area(r))
	fmt.Println("Perimeter: ", perimeter(r))

	r.Length *= 2
	r.Width *= 2
	fmt.Println("Double Area: ", area(r))
	fmt.Println("Double Perimeter: ", perimeter(r))
}
