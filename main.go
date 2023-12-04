package main

import "fmt"

func main() {
	var conferenceName string = "Aries Launch"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings [50]string //array
	var book2 []string      //slice
	bookings := []string{}

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

		if userTickets <= remainingTickets {
			book2 = append(book2, firstName+""+lastName) //slice
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("tickets left: %v \n", remainingTickets)
			fmt.Printf("Thank you for booking, %v. You will receive a confirmation email at %v. See you at the Con\n", firstName, email)
			bookings[0] = firstName + " " + lastName

		} else {
			fmt.Println("Please pick a valid number, we do not have that many tickets")
		}

	}

}
