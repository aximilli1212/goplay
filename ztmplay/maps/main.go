package main

import "fmt"

func main() {

	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 3
	shoppingList["milk"] = 35
	shoppingList["bread"] += 3
	shoppingList["eggs"] = 3
	shoppingList["cereal"] = 355

	shoppingList["eggs"] += 3

	fmt.Println(shoppingList)

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
