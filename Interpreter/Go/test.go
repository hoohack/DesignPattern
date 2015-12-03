package main

import (
  "./interpreter"
  "fmt"
)

func main() {
  statement := "3 * 2 * 4"
  calculator := &interpreter.Calculator{Symbol_arr: [5]string{"+", "-", "*", "/", "%"}}
  calculator.Init(statement)
  result := calculator.Compute()
  fmt.Printf("%s = %d\n", statement, result)
}
