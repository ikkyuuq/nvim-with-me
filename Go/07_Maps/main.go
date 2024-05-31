package main

import (
  "fmt"
  "strings"
)

func WordCount(s string) map[string]int {
  out := make(map[string]int)
  for _, word := range strings.Fields(s) {
    out[word]++
  }
  return out
}

func main() {
  result := make(map[string]int)
  result = WordCount("Hello World I am Kittipong Prasompong. World is beautiful with NVIM")
  fmt.Println(result)
}
