package main

import "fmt"

func variables() {

	const conferenceTickets int = 50 // con const no se puede reasignar la variable recordar
	var remainingTickets uint = 50   // tambien se puede escribir como value1 := "valor1"
	conferenceName := "Go Conference"

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	var firstName string // podes declarar el tipo de la variable por aparte o explicitamente en la variable
	var lastName string
	var email string
	var userTickets uint

	// asking for user input
	fmt.Scanln(firstName) // hace referencia al valor de la variable
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName) // hace referencia al espacio en memoria donde se encuentra el valor de la variable

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	// book ticket in system
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	/*
		Operation
	*/
	//relacionales
	fmt.Println("X == Y devuelve falso")
	fmt.Println("X != Y devuelve verdadero")
	fmt.Println("1 > 2 devuelve falso")
	fmt.Println("1 < 2 devuelve verdadero")
	fmt.Println("1 >= 1 devuelve verdadero")
	fmt.Println("2 <= 1 devuelve falso")

	//logicos
	fmt.Println("var A := true\nvar B := false\nvar X := 9\nvar Y := 5")

	fmt.Println("A && B resulta false // X > 0 && Y < 6 resulta true")
	fmt.Println("A || B resulta true // X < 0 || Y > 6 resulta false")
	fmt.Println("!A resulta false // !B resulta true")
}
