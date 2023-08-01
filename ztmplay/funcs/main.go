//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetMe(name string) {
	fmt.Println("Hello there:: ", name)
}

func returnMessage() string {
	return "Hello there message "
}

func printUs(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func returnNumber() int {
	return 100
}

func returnNums() (int, int) {
	return 100, 200
}

func main() {

	greetMe("Martin Luther")
	fmt.Println(returnMessage())
	fmt.Println(printUs(1, 2, 3))
	fmt.Println(returnNums())
	fmt.Println(returnNumber())
	fmt.Println(printUs(returnNumber(), returnNumber(), 100))

}
