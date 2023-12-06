package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Aries Launch"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	//	var bookings [50]string //array
	var book2 []string //slice
	//bookings := []string{} //method two of declaring slice

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have %v tickets left \n", remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var (
			firstName   string
			lastName    string
			userTickets uint
			email       string
		)

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets you want: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("sorry we do not have that many tickets")
			break
		}
		remainingTickets = remainingTickets - userTickets

		book2 = append(book2, firstName+""+lastName) //slice
		fmt.Printf("Thank you for booking, %v. You will receive a confirmation email at %v. See you at the Con\n", firstName, email)

		firstNames := []string{}
		for _, booking := range book2 {
			var names = strings.Fields(booking) //splits the string with white space as separator
			if len(names) > 0 {
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of ticket holders: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are booked out")
				break
			}

		}
	}

}
