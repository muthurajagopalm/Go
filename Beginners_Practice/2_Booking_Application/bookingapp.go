package main

import (
	"bookingapp/helper"
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			//function print firstNames//
			firstNames := getFirstNames()
			fmt.Printf("The first name of the users are %v\n", firstNames)

			if remainingTickets == 0 {
				//end the program
				println("Our Conference is booked out, Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email-id is incorrect - It doesn't contains @")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}

	}
}
func greetUser() {
	fmt.Println("Welcome to our conference!")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ar userTickets int
	//ask user for their name

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for the user//
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking the tickets in %v, You have booked %v tickets and the remaining tickets available is %v\n", firstName, lastName, conferenceName, userTickets, remainingTickets)
	fmt.Printf("These are all the bookings in our application %v\n", bookings)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########################")
	fmt.Printf("Sending tickets:\n %v  \nto email address %v\n", ticket, email)
	fmt.Println("###########################")
}
