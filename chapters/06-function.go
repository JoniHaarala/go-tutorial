package main

import (
	"fmt"
	//"strings"
)

/*
// packages level variables
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
*/

func functions() {
	/*
		tratamos siempre de evitar escribir todo el codigo en un solo lugar ya que estariamos creando un software monolitico.
		Ademas la logica del programa se vuelve cada vez mas dificil de manejar y testear.
		Por lo que descomponemos el programa general en partes como ya lo hice anteriormente, encapsulando el codigo en pequeños componentes donde cada uno tiene su propia funcionalidad y logica.

		Como pasarle parametros a una funcion: usando argumentos y definiendo de la siguiente manera:
			func funcName(arg1 type, arg2 type, ... ,argN type){}, donde type es el tipo de dato que recibe(int, string, bool, etc).
		Tambien puede retornar un valor mediante la palabra reservada RETURN.

		En GO tenes que declarar tanto el parametro del INPUT como el OUTPUT, con sus respectivos tipos de datos. Ejemplo:

			func funcName(arg1 string) string {		//el input se declara adentro de() y el output afuera
				var result string
				//some code here
				return result	//podemos retornar varios valores tambien
			}

		Tambien existen las packages level variable (variables a nivel de paquete) que son aquellas las cuales cualquier funcion que se encuentre
		dentro del paquete puede acceder a ellas, incluyendo el main y las demas funciones. Son como las variables globales del paquete (no de todo el programa ojo). Una cosa
		a tener en cuenta es que las variables a nivel de paquete no puede inicializarse con el asignador :=, por lo que siosi se una el VAR o CONST.
		Como buena practicar tratar de usar mas las variables LOCALES y no tanto las globales.
	*/

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

/*
func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
*/
