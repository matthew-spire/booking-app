package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to our %s booking application.\n", conferenceName)
	fmt.Printf("Get your tickets here to attend. The %s is limited to: %d people.\n", conferenceName, conferenceTickets)
	fmt.Printf("Currently there are %d still available.\n", remainingTickets)

	// The following creates an infinite for loop
	LOOP: for {
		var userFirstName string
		var userLastName string
		var emailAddress string
		var userTickets uint

		fmt.Print("What is your first name? ")
		fmt.Scanf("%s\n", &userFirstName)
		fmt.Print("What is your last name? ")
		fmt.Scanf("%s\n", &userLastName)
		fmt.Print("What is your email address? ")
		fmt.Scanf("%s\n", &emailAddress)
		fmt.Print("How many tickets would you like to purchase? ")
		fmt.Scanf("%d\n", &userTickets)

		validNameLength := len(userFirstName) >= 2 && len(userLastName) >= 2
		validEmail := strings.Contains(emailAddress, "@")
		validTickets := userTickets > 0 && userTickets <= remainingTickets

		if validTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, userFirstName + " " + userLastName)

			fmt.Printf("A booking of %d ticket(s) for %s %s to attend the %s is confirmed. A confirmation has been sent to %s. We look forward to seeing you!\n", userTickets, userFirstName, userLastName, conferenceName, emailAddress)
			fmt.Printf("The number of %s tickets left is: %d.\n", conferenceName, remainingTickets)
			fmt.Printf("The current %s bookings are: %v.\n", conferenceName, bookings)

			var firstNames []string
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				// var firstName = names[0]
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of the individuals who booked %s tickets are: %s.\n", conferenceName, firstNames)

			// var noTicketsRemaining bool = remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Printf("The %s is sold out! Please come again next year!", conferenceName)
				break
			}
		} else {
			fmt.Printf("There are insufficient %s tickets remaining to complete your request. You can book a maximum of %d tickets.\n", conferenceName, remainingTickets)
			goto LOOP
		}
	}
}
