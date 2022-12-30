package main

import "fmt"

func main (){

				var conferenceName = "Go Conference"
				const conferenceTickets = 50
				var remainingTickets = 50
				bookings := [] string{}

				fmt.Printf("Welcome to %v booking application \n", conferenceName)
				fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)
				fmt.Printf("Get your tickets here to attend \n")

				var firstName string
				var lastName string
				var userTickets int
				var email string

				for {
				fmt.Printf("Enter your First Name: \n")
				fmt.Scan(&firstName)


				fmt.Printf("Enter your Last Name: \n")
				fmt.Scan(&lastName)

				fmt.Printf("Enter number of tickets: \n")
				fmt.Scan(&userTickets)

				fmt.Printf("Enter your email here: \n")
				fmt.Scan(&email)

				remainingTickets = remainingTickets - userTickets
				bookings = append(bookings, firstName + " " + lastName)

				fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v \n", firstName, lastName, userTickets, email)
				fmt.Printf("%v Tickets remaining for %v\n", remainingTickets, conferenceName)
				fmt.Printf("These are all our bookings: %v\n", bookings)
				}

}