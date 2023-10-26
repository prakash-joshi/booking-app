package main

import "fmt"

func main() {
	var conferenceName string = "Golang Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("Welcome to %v Booking Application! \n", conferenceName)
	fmt.Printf("Only %v tickets are still available from %v tickets. \n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Print("Enter Your Fisrt Name : ")
	fmt.Scan(&firstName)

	fmt.Print("Enter Your Last Name : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter Your Email : ")
	fmt.Scan(&email)

	fmt.Print("Enter No of Tickets : ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole booking Array : %v.\n", bookings)
	fmt.Printf("The booking Array first value : %v.\n", bookings[0])
	fmt.Printf("The whole booking Array Type : %T.\n", bookings)
	fmt.Printf("The whole booking Array length : %v.\n", len(bookings))

	fmt.Printf("Thank You %v %v for booking %v tickts. You will receive the confirmation email at %v.\n", firstName, lastName, userTicket, email)
	fmt.Printf("Only %v tickets remaining for %v.\n", remainingTickets, conferenceName)
	// fmt.Println("")
	// fmt.Printf("conferenceName is %T. \n", conferenceName)
	// fmt.Printf("conferenceTickets is %T. \n", conferenceTickets)
	// fmt.Printf("remainingTickets is %T. \n", remainingTickets)
	// fmt.Printf("userName is %T. \n", userName)
	// fmt.Printf("userTicket is %v. \n", &userTicket)

}
