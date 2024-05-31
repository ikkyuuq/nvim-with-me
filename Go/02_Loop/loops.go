package main

import "fmt"

func main() {
	// For loop
	// just like a C pattern but without ()
	Str := ""
	for i := 0; i < 5; i++ {
		Str += "For"
	}
	fmt.Println(Str)

	// You want While loop?
	// Here is it
	for Str != "ForForForForForWhileWhileWhile" {
		Str += "While"
	}
	fmt.Println(Str)

	// Infinite loop? easy!
	// Just no condition
	// for {
	// 	Str += "Infinite"
	// }
	// Code below will never be executed because the loop above is Infinite
	// fmt.Println(Str)

	// Looping through array or slice with "range"
	// range returns index and value
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, value)
	}
	// If you don't need index, you can use _ (underscore) to ignore it
	for _, value := range arr {
		fmt.Println(value)
	}
}
