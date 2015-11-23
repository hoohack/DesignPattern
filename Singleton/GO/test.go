package main

import "fmt"
import "./singleton"

func main() {
  instance := singleton.GetInstance()
  another_instance := singleton.GetInstance()
  if instance != another_instance {
    fmt.Printf("not equal");
  }
}
