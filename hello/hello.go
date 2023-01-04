package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

/* To run type in console?
 * go run .
 * Add module dependency go mod edit -replace example.com/greetings=../greetings
 * Update module which will import that module go mod tidy
 */
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	//message, err := greetings.Hello("Gladys")

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	//message := greetings.Hello("Gladys")
	//fmt.Println(message)
	fmt.Println(messages)
}
