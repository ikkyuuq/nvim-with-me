package main

import (
  "fmt"
  "strconv"
)

func main() {
  s := "55"
  fmt.Println("Orinal String #1:", s)

  v, _ := strconv.Atoi(s)
  fmt.Println("String to Int:", v)

  ss := strconv.Itoa(v)
  fmt.Println("Int to String:", ss)
  
  fmt.Println("=====================================")

  str := "Hello, 世界"
  fmt.Println("Orinal String #2:", str)
  b := []byte(str)
  fmt.Println("String to Byte:", b)

  sss := string(b)
  fmt.Println("Byte to String:", sss)

}
