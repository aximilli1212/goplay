package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is Clean")
		} else {
			fmt.Println(room.name, "is Dirty")
		}
	}
}

func main() {

	hotelRooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Salle"},
		{name: "Reception"},
	}

	checkCleanliness(hotelRooms)

	fmt.Println("Hold on... Performing cleaning...")
	hotelRooms[0].cleaned = true
	hotelRooms[2].cleaned = true

	checkCleanliness(hotelRooms)
}
