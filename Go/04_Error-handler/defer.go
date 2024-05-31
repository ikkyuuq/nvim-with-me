package main

import "fmt"

func main() {
  // Defer Statement
  // A defer statement will execute the function call at the end of the enclosing function
  // Defer is commonly used to ensure that a file is closed immediately after it has been opened
  // Defer is also used to close a connection to a database

  // Most clear example of defer
  b()
  // Output:
  // entering: b // Trace will be called first btw, un ["b"]
  // in b // From print in func b
  // entering: a // Trace will be called first btw, un ["a", "b"]
  // in a // From print in func a
  // So the program is done executing the function 
  // un ["a", "b"] // un will be called before the end of program
  // leaving: a // un will be called before the end of program
  // leaving: b // un will be called before the end of program

  // As you can see a is called last and b is called first but the output is in reverse order
  // This is because defer is called at the end of the function
  // called LIFO (Last In First Out)
}

func trace(s string) string {
  fmt.Println("entering:", s)
  return s
}

func un(s string) {
  fmt.Println("leaving:", s)
}

func a() {
  defer un(trace("a")) // Trace will be called first, then un will be called before the end of program
  fmt.Println("in a")
}

func b() {
  defer un(trace("b")) // Trace will be called first, then un will be called before the end of program
  fmt.Println("in b")
  a()
}
