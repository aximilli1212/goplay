package main

import "fmt"

func checkMaturity(userAge int) string {

	switch age := userAge; {
	case age < 1:
		return "newborn"

	case age >= 1 && age <= 3:
		return "toddler"

	case age >= 4 && age <= 12:
		return "child"

	case age >= 13 && age <= 17:
		return "Teenager"

	case age >= 18:
		return "Adult"

	default:
		return "Alien"

	}
}

func main() {
	fmt.Println(checkMaturity(2))

}
