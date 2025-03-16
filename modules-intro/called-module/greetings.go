package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello exported name: a function whose name starts with a capital letter can be called by a function outside its package
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}

	// := operator is a shortcut for declaring and initializing variable in the same line
	// the value on the right side of the operator is used to decide the type of the variable
	message := fmt.Sprintf(randomGreeting(), name)
	// message := fmt.Sprint(randomGreeting())

	/*
		The long way:
			var message string
			string = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/

	return message, nil
}

// Hellos returns a map that associates names with greetings
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomGreeting() string {
	// we have initialized a slice, which is a dynamic array (array list of Go)
	greetings := []string{
		"Hi, %v. Welcome!",
		"Great to have you, %v!",
		"Have a lovely day, %v!",
	}

	// from the math/rand package we use the Intn function and pass the length of slice greetings to limit the range of rand
	return greetings[rand.Intn(len(greetings))]
}
