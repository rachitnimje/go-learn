package greetings

import "fmt"

// exported name: a function whose name starts with a capital letter can be called by a function outside its package
func Hello(name string) string {
	// := operator is a shortcut for declaring and initializing variable in the same line
	// the value on the right side of the operator is used to decide the type of the variable
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	/*	
		The long way:
			var message string
			string = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/

	return message
}