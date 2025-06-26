package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type userData struct {
	firstName     string
	lastName      string
	email         string
	contactNumber int
	userTickets   int
}

var wg sync.WaitGroup

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", conferenceName, " booking application")
	fmt.Println("We have total of ", conferenceTickets, "tickets remaining", remainingTickets)
	fmt.Println("Get your tickets here to attend")
	booking := []userData{}

	var firstName string
	var lastName string
	var email string
	var contactNumber int
	var userTickets int
	fmt.Println("\nEnter Your firstName:")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your lastName:")
	fmt.Scan(&lastName)
	fmt.Println("Eneter your email  :")
	fmt.Scan(&email)
	fmt.Println("The conformation mail will send your email kindly check it ")
	fmt.Println("Enter your contact Number:")
	fmt.Scan(&contactNumber)
	fmt.Println("Eneter your userTickets  :")
	fmt.Scan(&userTickets)

	isValidName := len(firstName) >= 2 && len(lastName) >= 1
	isValidEmail := strings.Contains(email, "@")
	length := len(strconv.Itoa(contactNumber))
	isValidContactNumber := length <= 10
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	if isValidEmail && isValidName && isValidTicketNumber && isValidContactNumber {
		remainingTickets = remainingTickets - userTickets
		newBookings := userData{
			firstName:   firstName,
			lastName:    lastName,
			email:       email,
			userTickets: userTickets,
		}
		booking = append(booking, newBookings)
		fmt.Printf("\n Thank you %v %v for booking %v tickets!\n", firstName, lastName, userTickets)
		fmt.Printf("Confirmation will be sent to %v \n", email)
		fmt.Printf("We will contact you on %v\n", contactNumber)
		fmt.Printf("Remaining tickets: %v\n", remainingTickets)
		fmt.Printf("Number of Bookings: %v\n", userTickets)
		fmt.Printf("People who booked : %v\n ", getFirstName(booking))
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		if remainingTickets == 0 {
			fmt.Println("All tickets are booked. Thankyou")
		}
	} else {
		fmt.Println("\nInvalid input:")
		if !isValidName {
			fmt.Println("-Name if not valid")
		}
		if !isValidTicketNumber {
			fmt.Println("Invalid TicketNumber")
		}
		if !isValidEmail {
			fmt.Println("Invalid Email")
		}
		if isValidContactNumber {
			fmt.Println("Invelid Contact Number:")
		}
	}
	wg.Wait()
}
func getFirstName(booking []userData) []string {
	firstNames := []string{}
	for _, booking := range booking {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}
func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %V", userTickets, firstName, lastName, email)
	fmt.Println("***********************************")
	fmt.Printf("sending Tickets :\n %v \n to email address %v\n", ticket, email)
	fmt.Println("**********************************")
	wg.Done()
}
