package main

import "fmt"

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

	fmt.Println(min1, min2, min3) //NB: all the declared consts are assigned the first assigned vlaue: 400 by default
}
