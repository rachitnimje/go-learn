// we declare a main package
// a package is a way to group functions and its made up of all the files in that directory
// Go code is grouped into packages, and packages are grouped into modules. 
package main

/* 
	fmt is a standard go library which provides formatting for text including printing to console
	when ytou import package from another module, you run "go mod tidy" to locate and download the module
	this action will add the module name and version under the "require" section in the go.mod file
	this will also create a go.sum file which is used for authenticating the module 
*/
import (
	"fmt"
	"rsc.io/quote"
)	

// the starting point of execution is main package and the main function inside it
func main() {
	fmt.Println("hello gophers!")
	fmt.Println(quote.Go())
}
