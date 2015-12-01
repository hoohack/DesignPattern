package main

import "./flyweight"
import "fmt"

func main() {
  mgr := flyweight.GetSInstance()
  mgr.Login("hhq")
  mgr.Login("tom")
  fmt.Println(mgr.HasPermit("hhq", "money", "look"))
  fmt.Println(mgr.HasPermit("tom", "money", "look"))
}
