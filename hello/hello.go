package main
// In Go, code executed as an application must be in a main package
import (
	"fmt"
	"log"


	"example.com/greetings"
)

func main() {
 	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

    // A slice of names.
    names := []string{"Daniel", "JK", "Debo"}

	// get greetings from greetings module and use it
	message, err := greetings.Hellos(names)
	
	// If an error was returned, print it to the console and exit the program.
    if err != nil {
        log.Fatal(err)
    }

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}