package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

type Product struct {
	Price int
	Name  string
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

func totalCost(prods [4]Product) int {
	sum := 0
	for i := 0; i < len(prods); i++ {
		prod := prods[i]
		sum += prod.Price
	}

	return sum
}

func main() {

	hotelRooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Salle"},
		{name: "Reception"},
	}

	ShoppingList := [4]Product{
		{500, "KeySoap"},
		{500, "Mango Soap"},
		{500, "Rain Soap"},
	}

	checkCleanliness(hotelRooms)

	fmt.Println("Hold on... Performing cleaning...")
	hotelRooms[0].cleaned = true
	hotelRooms[2].cleaned = true

	checkCleanliness(hotelRooms)

	fmt.Println("Last item:", ShoppingList[2])
	fmt.Println("Lenght of items:", len(ShoppingList))
	fmt.Println("Last item:", ShoppingList[2])
	fmt.Println("Cost of items", totalCost(ShoppingList))

	ShoppingList[3].Name = "MOngoosa"
	ShoppingList[3].Price = 566

	fmt.Println(ShoppingList)
}
