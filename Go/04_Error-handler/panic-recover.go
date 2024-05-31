package main

import "fmt"

func main() {
	// Panic
	// A panic is a situation when a program cannot continue to run it will stop the program and defer will be called
	// Panic is a way to stop the program when something goes wrong
	// Recover is a way to handle the panic
	// r := recover()
	// r will be nil if there is no panic
	// r will be the value of panic if there is a Panic
	// Recover is only useful when it is called inside the defer function
	// Recover will only work if it is called inside the defer function

	f()
	fmt.Println("Returned normally from f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r) // Called after defer in g() end
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") // Never executed because of the panic in g make the program stop and defer will be called
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i)) // This will stop the program and defer in g() will be called
	}
	defer fmt.Println("Defer in g", i) // This will be called before the panic
	fmt.Println("Printing in g", i)    // Normal print until the panic
	g(i + 1)
}
