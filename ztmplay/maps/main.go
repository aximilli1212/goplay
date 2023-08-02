package main

import "fmt"

func printServerDetails() {

}

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {

	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	shoppingList := make(map[string]int)
	serverList := make(map[string]int)

	for _, server := range servers {
		serverList[server] = Offline
	}

	shoppingList["eggs"] = 3
	shoppingList["milk"] = 35
	shoppingList["bread"] += 3
	shoppingList["eggs"] = 3
	shoppingList["cereal"] = 355

	shoppingList["eggs"] += 3

	fmt.Println(shoppingList)
	fmt.Println(serverList)

	delete(shoppingList, "milk") // delete an item from the map

	cereal, ok := shoppingList["cereal"]
	fmt.Println("Need Cereal?")
	if !ok {
		fmt.Println("Nope we dont have that.")
	} else {
		fmt.Printf("Yup we got %d  boxes \n", cereal)
	}

	for key, val := range shoppingList {
		fmt.Println(key, val)
	}
}
