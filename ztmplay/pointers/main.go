package main

import (
	"fmt"
)

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("counter:", counter)
	fmt.Println("counter22:", counter.hits)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
	fmt.Println("0ld string:", *old)

}

type Storage struct {
	Name string
	Tag  bool
}

func printSlice(slice []*Storage) {
	for _, value := range slice {
		fmt.Println(*value)
	}
}

func checkout(tags []*Storage) {
	for _, value := range tags {
		value.Tag = false
	}
}
func main() {

	counter := Counter{}

	hello := "Hello"
	//world := "World"

	replace(&hello, "HoiI", &counter)

	//--Requirements:
	//* Create a structure to store items and their security tag state
	//  - Security tags have two states: active (true) and inactive (false)
	//* Create functions to activate and deactivate security tags using pointers
	//* Create a checkout() function which can deactivate all tags in a slice
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array

	secStore := []*Storage{
		{"Pantry", true},
		{"Janitorial", true},
		{"Nursery", true},
		{"Office", true},
	}

	//  - Deactivate any one security tag in the array/slice
	secStore[1].Tag = false
	printSlice(secStore)

	//  - Call the checkout() function to deactivate all tags
	checkout(secStore)
	fmt.Println("-------")
	printSlice(secStore)

	//  - Print out the array/slice after each change
}
