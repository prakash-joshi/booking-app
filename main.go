package main

import "fmt"

func main() {
	var conferenceName string = "Golang Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v Booking Application! \n", conferenceName)
	fmt.Printf("Only %v tickets are still available from %v tickets. \n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTicket int

	userName = "Prakash"
	userTicket = 23
	fmt.Printf("User %v has ticket no : %v. \n", userName, userTicket)

	fmt.Println("")

	fmt.Printf("conferenceName is %T. \n", conferenceName)
	fmt.Printf("conferenceTickets is %T. \n", conferenceTickets)
	fmt.Printf("remainingTickets is %T. \n", remainingTickets)
	fmt.Printf("userName is %T. \n", userName)
	fmt.Printf("userTicket is %T. \n", userTicket)

}
