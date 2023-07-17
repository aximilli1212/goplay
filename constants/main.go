package main

import (
	"fmt"
	"math"
)

func main() {
	const days = 45
	fmt.Println("numbered Days are: ", days)
	//Note that Variable errors are caught at runtime but contant error e.g:
	//Dividing by 0 are caught at compile time: right in the IDE

	const (
		min1 = 400
		min2
		min3
	)

	fmt.Println(min1, min2, min3) //NB: all the declared consts are assigned the
	// first assigned vlaue: 400 by default
	infVal := math.Pow(3, 4)
	fmt.Println(infVal)

	//Constant Rules
	//1. You cannot change a constant
	//2. You cannot initialize a const at runtime
	//const power = math.Pow(2,3) will throw an error
	//3. You cannot use a constant to initialize a variable
	//4. You can initialize a const with a builtIn func eg: len

}
