package main

import (
	"Ticky/helper"
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName string = "Aries Launch"
var remainingTickets uint = 50
var book2 []string //slice
//bookings := []string{} //method two of declaring slice

func main() {

	greetings()

	for {

		//getting user input
		firstName, lastName, email, userTickets := getUserInput()

		//validate user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//conditionals for response
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket()

			fNames := getFNames()
			fmt.Printf("The first names of bookings are: %v\n", fNames)

			//logic to handle no tickets scenario
			if remainingTickets == 0 {
				fmt.Println("Sorry, we are booked out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please fill in name field correctly(cannot be less than 3 characters)")
			}
			if !isValidEmail {
				fmt.Println("Sorry your email is invalid")
			}
			if !isValidTicketNUmber {
				fmt.Println("Please input a valid ticket number")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have %v tickets left \n", remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFNames() {
	//formatting user details
	firstNames := []string{}
	for _, booking := range book2 {
		var names = strings.Fields(booking) //splits the string with white space as separator
		if len(names) > 0 {
			firstNames = append(firstNames, names[0])
		}
	}
	return firstNames
}

func getUserInput() (string, string, uint, string) {
	var (
		firstName   string
		lastName    string
		userTickets uint
		email       string
	)

	//Request User input using pointers to await/delay
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets, email

}

func bookTicket(userTickets uint, book2 []string, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets
	book2 = append(book2, firstName+""+lastName) //slice

	fmt.Printf("Thank you for booking, %v. You will receive a confirmation email at %v. See you at the Con\n", firstName, email)

}
