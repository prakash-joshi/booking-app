package main

import "fmt"

func main() {
	var conferenceName string = "Golang Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v Booking Application! \n", conferenceName)
	fmt.Printf("Only %v tickets are still available from %v tickets. \n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTicket int
	fmt.Print("Enter Your Fisrt Name : ")
	fmt.Scan(&firstName)

	fmt.Print("Enter Your Last Name : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter Your Email : ")
	fmt.Scan(&email)

	fmt.Print("Enter No of Tickets : ")
	fmt.Scan(&userTicket)

	fmt.Printf("Thank You %v %v for booking %v tickts. You will receive the confirmation email at %v.\n", firstName, lastName, userTicket, email)

	// fmt.Println("")
	// fmt.Printf("conferenceName is %T. \n", conferenceName)
	// fmt.Printf("conferenceTickets is %T. \n", conferenceTickets)
	// fmt.Printf("remainingTickets is %T. \n", remainingTickets)
	// fmt.Printf("userName is %T. \n", userName)
	// fmt.Printf("userTicket is %v. \n", &userTicket)

}
