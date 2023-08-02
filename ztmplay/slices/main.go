package main

import (
	"fmt"
)

type Part string

func showPart(line []Part) {
	for i, part := range line {
		fmt.Println(part, i)
	}
}

func main() {

	//* Using a slice, create an assembly line that contains type Part
	//* Create a function to print out the contents of the assembly line
	//* Perform the following:
	//  - Create an assembly line having any three parts
	assemblyLine := []Part{"Screws", "Bolts", "Nuts"}
	fmt.Println("Original Assembly:", assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Wing", "Steering")
	fmt.Println("Updated Assembly: ", assemblyLine)

	//  - Slice the assembly line, so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println(assemblyLine)
	//  - Print out the contents of the assembly line at each step
	//--Notes:
	//* Your program output should list 3 parts, then 5 parts, then 2 parts

	//You can iterate over a slice using range
	showPart(assemblyLine)

}
