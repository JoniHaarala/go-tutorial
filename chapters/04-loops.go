package main

import (
	"fmt"
	"strings"
)

func loops() {
	/*
		In GO specifically, loops are simplified, you only have the "for" loop
		not "while", "do-while", "for-each"
	*/

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	// infinite loop (same as while)
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// asking for user input
		fmt.Println("Enter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		// book ticket in system
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		/*
			conditional loop (as a normal for)
		*/
		// ej: print only first names
		firstNames := []string{}

		// index: es el indice. En caso de no usarlo colocamos un _ en su lugar; value: es el valor dentro del array
		for _, booking := range bookings { //range iterates over elements in different structures, not only arrays.
			var names = strings.Fields(booking) //igual al array.split en js o python
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names %v\n", firstNames)
	}
}
