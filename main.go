package main

import "fmt"

func main() {
	var conferenceName = "Golang Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v Booking Application! \n", conferenceName)
	fmt.Printf("Only %v tickets are still available from %v tickets. \n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend.")

}
