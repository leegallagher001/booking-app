package main

import "fmt"

func main() {
	var conferenceName = "Go Conference 2026" // no need to specify data type here, as Go can infer the type as a value has been applied during variable declaration
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string // array has specified number of elements, slice does not so is more dynamic

	fmt.Printf("conferenceName is of type %T, conferenceTickets is of type %T and remainingTickets is of type %T\n", conferenceName, conferenceTickets, remainingTickets) // %T is format specifier to print variable type

    fmt.Printf("Welcome to the %v booking application!\n", conferenceName) // can be done using Println and string concatenation as well
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string // variable type specified as no value given during declaration
	var lastName string
	var emailAddress string
	var userTickets uint
	// ask user for their name
	fmt.Println("Please provide your first name: ")
	fmt.Scan(&firstName) // & is an example of a pointer, which is a variable that stores the memory address of another variable. In this case, it is used to store the memory address of the firstName variable, so that the Scan function can modify its value directly.
    // scan function reads input and stores it in variable found at the memory address provided by the pointer. It allows the program to capture user input and assign it to the specified variable.
	fmt.Println("Please provide your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please provide your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName) // adds a new booking element to the slice, which is a dynamic array that can grow or shrink in size as needed. The append function is used to add the new booking to the end of the slice.

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))
	fmt.Printf("Slice pointer: %v\n", &bookings[0])

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v of %v tickets remaining for %v\n", remainingTickets, conferenceTickets, conferenceName)

	fmt.Printf("These are all of the bookings so far: %v\n", bookings)
}

// Video - Golang Tutorial for Beginners | Full Go Course - 1:11:10