package main

import (
	"Ticky/helper"
	"fmt"
)

const conferenceTickets uint = 50

var conferenceName string = "Aries Launch"
var remainingTickets uint = 50
var book2 = make([]UserData, 0)

//var book2 = make([]map[string]string, 0) //slice
//bookings := []string{} //method two of declaring slice

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.waitGroup{}

func main() {

	greetings()

	for {

		//getting user input
		firstName, lastName, email, userTickets := getUserInput()

		//validate user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//conditionals for response
		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)


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
		wg.Wait()
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
		firstNames = append(firstNames, booking.firstName)

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

	//map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	/*userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	//converting uint to string as maps only take one datatype
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)*/

	book2 = append(book2, userData)

	fmt.Printf("Thank you for booking, %v. You will receive a confirmation email at %v. See you at the Con\n", firstName, email)

	func sendTicket(userTickets uint, firstName string, lastName string, email string) {
		time.Sleep(10 * time.Second)
		var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
		fmt.Println("#################")
		fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
		fmt.Println("#################")
		wg.Done()


	}

}
