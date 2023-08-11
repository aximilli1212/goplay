package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {

	myLot := ParkingLot{spaces: make([]Space, 5)}

	fmt.Println("initial:", myLot)

	myLot.occupySpace(1)
	occupySpace(&myLot, 2)
	fmt.Println("After occupied:", myLot)
	myLot.vacateSpace(4)
	fmt.Println("After vacate 3:", myLot)
}
