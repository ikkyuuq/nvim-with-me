package main

import "fmt"

// Delaration of Sqrt function
// It takes a float64 as parameter and returns a float64

// func <function_name>(<parameter_name> <parameter_type>) <return_type>

// Sqrt is a function to calculate square root of x
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
