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
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])

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
