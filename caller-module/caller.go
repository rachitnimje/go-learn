package main

import (
	"fmt"
	"halogen/called-module"
)

// in this module, I want to call a package from another module (greetings)
// for this firstly, we need to import the package in the import tag
// secondly, we need to set the dependency location of halogen/greetings such that it takes the module from the local file system
// for this, we run "go mod edit -replace <name of the module>=<local path for the module>"
// it will look like "go mod edit -replace halogen/greetings=../greetings"
// then we need to run "go mod tidy" so that it can include the external module 
func main() {
	// var message string
	message := greetings.Hello("Rachit") 
	fmt.Println(message)
}