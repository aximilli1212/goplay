package main

import "log"

//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

type Player struct {
	Health    uint
	Energy    uint
	MaxHealth uint
	MaxEnergy uint
	Name      string
}

func (player *Player) updateHealth(val uint) {
	player.Health += val

	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}
}

func (player *Player) updateEnergy(val uint) {
	player.Energy = val
}

func main() {

	myPlayers := []Player{
		{Health: 61, Energy: 85, Name: "Pablo Escobar"},
		{Health: 40, Energy: 85, Name: "Shanning Nunexa"},
		{Health: 20, Energy: 35, Name: "Mondixa Plolica"},
	}

	log.Println("initial player stat", myPlayers)
	myPlayers[2].updateEnergy(400)
	log.Println("After second player energy update", myPlayers)
	myPlayers[1].updateHealth(800)
	log.Println("after first player health update", myPlayers)

}
