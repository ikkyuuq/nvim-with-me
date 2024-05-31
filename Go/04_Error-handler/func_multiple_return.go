package main

import (
	"fmt"
	"log"
)

func main() {

	// This is a simple example of multiple return
	result, err := Encode("Hello")
	fmt.Println(result, err)

	result, err = EncodeWithDefer("Hello With Defer")
	fmt.Println(result, err)

}

func Encode(s string) (res string, err error) {
	// return s, nil
	// This will return the value of s and nil
	// But this is more readable
	// error != nil is more understandable than retrun nil
	res = s
	if err != nil {
		log.Fatal(err) // Exit with status 2
	}
	return res, err
	// BTW, this migth be a good practice
	// but it also can use defer statment to handle the error
}

func EncodeWithDefer(s string) (res string, err error) {
	res = s
	if err != nil {
		return "", err // If there is no error, return empty string and defer won't be called
	}
	// This will call the errorFunc() before the return
	// This is useful to handle the function that need to be called before the return
	// Like closing the file, closing the connection, etc
	defer deferFunc()

	return res, err
}

func deferFunc() {
	fmt.Println("Going to return")
}
