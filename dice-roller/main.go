package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dice, sides := 2, 12
	rolls := 1

	for r := 1; r <= rolls; r++ {
		fmt.Println("Rolling dice")
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Rolle #", r, "Dice #", d, "Rolled:", rolled)
		}
		fmt.Println("Total rolled:", sum)

		switch sum := sum; {
		case sum == 2:
			fmt.Println("You rolled a snake eyes!")
		case sum == 7:
			fmt.Println("You rolled a lucky seven!")
		case sum%2 == 0:
			fmt.Println("You rolled an even number!")
		case sum%2 != 0:
			fmt.Println("You rolled an odd number!")
		}
	}
}
