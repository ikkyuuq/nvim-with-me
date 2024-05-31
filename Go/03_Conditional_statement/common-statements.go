package main

import "fmt"

func main() {
  // Basic if statement
  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  // If statement with Initialization statement
  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  }

  // Flow control statement
  //// continue to next iteration
  //// break out of the loop
  //// return from the function
  //// goto to another part of the program

  // Switch statement
  i := 2
  fmt.Print("Write ", i, " as ")
  switch i {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3:
      fmt.Println("three")
  }

  // Type switch statement
  var t interface{}
  t = func (i int) string {return "hello"}
  switch t := t.(type) {
    case int:
      fmt.Println("t is an integer")
    case string:
      fmt.Println("t is a string")
    default:
      fmt.Printf("t is of type %T\n", t)
  }
}
