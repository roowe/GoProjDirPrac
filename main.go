package main

import (
  "./arithmetic"
  "fmt"
)

func main() {
  var a int
  var b int
  fmt.Scan(&a, &b)
  fmt.Printf("Added to %v\n", arithmetic.Add(a, b))
}
