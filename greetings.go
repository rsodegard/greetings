package main

import "fmt"

//Hello reutnrs a greeting for th enamed person.
fun Hello(name string) string {
	// Return a greeting that embeds the name in the message.
	message := fmt.Sprintf("hi, %v. Welcome!", name)
	return message
}