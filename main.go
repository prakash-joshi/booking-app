package main

import "fmt"

func main() {
	var conferenceName = "Golang Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "Booking Application!")
	fmt.Println("Only", remainingTickets, "are still available from", conferenceTickets, "tickets.")
	fmt.Println("Get your tickets here to attend.")

}
