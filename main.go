package main

import "fmt"

func main() {
	var conferenceName = "Go Conference 2026" // no need to specify data type here, as Go can infer the type as a value has been applied during variable declaration
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceName is of type %T, conferenceTickets is of type %T and remainingTickets is of type %T\n", conferenceName, conferenceTickets, remainingTickets) // %T is format specifier to print variable type

    fmt.Printf("Welcome to the %v booking application!\n", conferenceName) // can be done using Println and string concatenation as well
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string // variable type specified as no value given during declaration
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets!\n", userName, userTickets)
}

// Video - Golang Tutorial for Beginners | Full Go Course - 42:00