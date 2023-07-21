package main

import "fmt"

func main() {
	name := "Akumia"
	fmt.Println("Hello playground, I am: ", name)

	//Note that fmt.Println() can also accept expressions
	a, b := 4, 6
	fmt.Println("Sum: ", a+b, "Mean Value: ", (a+b)/2)

	//Note that there is fmt.Printf() which allows you to interpolate with placeholders
	fmt.Printf("x is %d , y is %f, stringa is: %s", 5, 4.5, "shook")

	//%v - Replace by any type
	//%T - Show me the type :: typeOf()
	//%f - Replace with Fload
	//$.3f - Convert float to show 3 decimal point

}
