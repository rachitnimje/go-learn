package greetings

import (
	"testing"
	"regexp"
)

// these command will run all the funtions starting with "Test" in the file ending with "_test.go"
// to execute a test, run the command "go test"
// to list all the tests and their results, run "go test -v"
// test functions take a pointer to the testing package's testing.T type as parameter

// this is a test function which calls the Hello function with a valid name and checks if the response contains the passed name
// if the response does not include the name which we passed in, we use the t parameter's Errorf method to print a message
func TestHelloName(t *testing.T) {
	name := "Rachit"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Rachit")

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Rachit") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// this is a test function which calls the Hello function with an empty string and checks if the response is an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}