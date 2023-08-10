package main

import "log"

type Coordinate struct {
	X, Y int
}

func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y
}

func main() {

	mycoord := Coordinate{5, 5}
	mycoord.shiftBy(1, 2)

	log.Println(mycoord)

}
