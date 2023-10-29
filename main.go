package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Golang Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v Booking Application! \n", conferenceName)
	fmt.Printf("Only %v tickets are still available from %v tickets. \n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend.")
	for {
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Print("Enter Your First Name : ")
		fmt.Scan(&firstName)
		fmt.Print("Enter Your Last Name : ")
		fmt.Scan(&lastName)
		fmt.Print("Enter Your Email : ")
		fmt.Scan(&email)
		fmt.Print("Enter No of Tickets : ")
		fmt.Scan(&userTicket)

		isValidName := len(firstName) > 1 && len(lastName) > 1
		isVAlidEmail := len(email) > 6 && strings.Contains(email, "@")
		isValidTicketCount := userTicket > 0 && userTicket <= remainingTickets

		if isValidName && isVAlidEmail && isValidTicketCount {
			remainingTickets = remainingTickets - userTicket
			bookings = append(bookings, firstName+" "+lastName) ///similar too push for array in js

			fmt.Printf("Thank You %v %v for booking %v tickts. You will receive the confirmation email at %v.\n", firstName, lastName, userTicket, email)
			fmt.Printf("Only %v tickets remaining for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, bookingElement := range bookings {
				var name = strings.Fields(bookingElement)
				firstNames = append(firstNames, name[0])
			}

			fmt.Printf("These are the Names of our bookings : %v. \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry the confrence is sold out. Come back next Year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The First Name or Last Name is too short.")
			}
			if !isVAlidEmail {
				fmt.Println("The email id is incorrect.")
			}
			if !isValidTicketCount {
				fmt.Println("The ticket count you entered is invalid.")
			}
			fmt.Println("Please try again.")
			continue
		}
	}

	// fmt.Printf("The whole booking Slice : %v.\n", bookings)
	// fmt.Printf("The booking Slice first value : %v.\n", bookings[0])
	// fmt.Printf("The whole booking Slice Type : %T.\n", bookings)
	// fmt.Printf("The whole booking Slice length : %v.\n", len(bookings))
	// fmt.Println("")
	// fmt.Printf("conferenceName is %T. \n", conferenceName)
	// fmt.Printf("conferenceTickets is %T. \n", conferenceTickets)
	// fmt.Printf("remainingTickets is %T. \n", remainingTickets)
	// fmt.Printf("userName is %T. \n", userName)
	// fmt.Printf("userTicket is %v. \n", &userTicket)

}
