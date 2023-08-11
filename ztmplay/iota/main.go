package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs

	default:
		panic("unhandled exception")

	}
}

func main() {
	add := Operation(Add)
	sub := Operation(Subtract)
	mul := Operation(Multiply)
	div := Operation(Divide)
	add.calculate(39, 34)

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
