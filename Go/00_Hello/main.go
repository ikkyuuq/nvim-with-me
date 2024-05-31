// Declaring package is a way to group functions in same directory
package main

// Importing fmt package
// fmt package implements formatted I/O with functions analogous to C's printf and scanf
import "fmt"

// Importing quote package (External package)
// :!go get rsc.io/qoute (To install the package from the internet)
import "rsc.io/quote"

// main function is the entry point of the program
func main() {
	// Println function is used to print the message to the standard output
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}

// go get: Primarily used to manage (add, update, or remove) dependencies.
// go mod tidy: Used to clean up the go.mod and go.sum files, by removing any dependencies that are no longer required.
