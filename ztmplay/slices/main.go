package main

import "fmt"

type Part string

func main() {

	//* Using a slice, create an assembly line that contains type Part
	//* Create a function to print out the contents of the assembly line
	//* Perform the following:
	//  - Create an assembly line having any three parts
	assemblyLine := []Part{"Screws", "Bolts", "Nuts"}
	fmt.Println("Original Assembly:", assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Wing", "Steering Wheel")
	fmt.Println("Updated Assembly: ", assemblyLine)

	//  - Slice the assembly line, so it contains only the two new parts
	//  - Print out the contents of the assembly line at each step
	//--Notes:
	//* Your program output should list 3 parts, then 5 parts, then 2 parts

}
