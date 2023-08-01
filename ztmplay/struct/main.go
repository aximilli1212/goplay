package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat  Passenger
	MiddleSeat Passenger
	BackSeat   Passenger
	Driver     string
}

func main() {

	casey := Passenger{"Casey", 123, true}

	fmt.Println(casey)

	var (
		bill     = Passenger{Name: "Bill", TicketNumber: 3}
		sharlene = Passenger{TicketNumber: 2, Name: "Sharlene"}
	)

	fmt.Println(bill)
	fmt.Println(sharlene)

	bill.Boarded = true
	sharlene.Boarded = true

	var heidi Passenger
	heidi.Name = "Heidi"
	fmt.Println(heidi)

	bus := Bus{FrontSeat: heidi, MiddleSeat: bill, BackSeat: sharlene, Driver: "Casey"}

	fmt.Println(bus.FrontSeat)
	fmt.Println(bus.MiddleSeat.Boarded)

}
