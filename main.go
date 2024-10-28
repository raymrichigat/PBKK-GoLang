package main

import "fmt"

func main() {
		var conferenceName = "Go Conference"
		const conferenceTickets = 100
		var remainingTickets = 50

		fmt.Printf("Welcome to %v Ticketing System\n", conferenceName)
		fmt.Printf("We have %v tickets available for the conference\n", conferenceTickets)
		fmt.Printf("We have %v tickets remaining for the conference\n", remainingTickets)
		fmt.Println("Please enter your name: ")

		
}
