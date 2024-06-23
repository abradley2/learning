package main

import (
	"errors"
	"fmt"
	"strings"
)

// can only have one "main" function in any program
func main() {
	// value can be reset
	// := is infered variable can not declair types
	conferenceName := "Go Conference"
	// constants. Values are locked
	const conferenceTickets = 50
	// set stating value
	var remainingTickets uint = conferenceTickets

	var userNumberOfTickets uint
	// slice
	bookings := []string{}
	// var bookings []string
	// var bookings  = []string{}
	firstNames := []string{}

	// fmt.Printf("conferenceTickets type: %T, remainingTickets type: %T\n", conferenceTickets, remainingTickets)
	for {

		// there are different types of placeholders for different types and formatting
		// fmt.Println("Hello, would you like to play a game?")

		greating(conferenceName, conferenceTickets, remainingTickets)
		// ask user for name
		firstName, errFirstName := getName("What is your first name?")
		lastName, errLastName := getName("What is your last name?")
		if errFirstName != nil || errLastName != nil {
			fmt.Println("Names must be at least 2 characters. Please restart.")
			continue
		}
		// // takes a pointer as an argument and assigns the user input to
		// // that variable
		// fmt.Scan(&lastName)

		fmt.Println("How many tickets would you like to buy?")
		fmt.Scan(&userNumberOfTickets)
		if userNumberOfTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining. You will need to reduce the number of tickets you're buying \n")
			continue
		}

		// update remainingTickets
		remainingTickets = remainingTickets - userNumberOfTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("%v, you've bought %v tickets!\n", firstName+" "+lastName, userNumberOfTickets)
		fmt.Printf("%v remain\n", remainingTickets)
		fmt.Printf("All our bookings: %v\n", bookings)
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The full list is: %v/n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is sold out! Better luck next time!")
			break
		}
	}
}

func greating(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Println("Welcome to our program")
	fmt.Println("Get your tickets to attend", conferenceName, "now!")
	fmt.Printf("Total of %v for this conference\n", conferenceTickets)
	fmt.Println(remainingTickets, "left!")
}

func getName(prompt string) (string, error) {
	attempts := 0
	var firstName string
	fmt.Println(prompt)
	// takes a pointer as an argument and assigns the user input to
	// that variable
	fmt.Scan(&firstName)
	isValidName := len(firstName) >= 2
	for !isValidName && attempts < 3 {
		println("Names must be at least 2 characters long.")
		println("Please enter a name longer then 2 characters.")
		attempts++
		fmt.Scan(&firstName)
	}
	if !isValidName {
		return firstName, errors.New("Max number of attempts to enter a vailid name reached")
	}
	return firstName, nil
}
