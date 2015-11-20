package main

import (
  "fmt"
  "./factorymethod"
)

func main() {
  iphoneFactory := &factorymethod.IPhoneFactory{}
  androidFactory := &factorymethod.AndroidFactory{}
  otherFactory := &factorymethod.OtherFactory{}

  iphone := iphoneFactory.CreatePhone("6")
  fmt.Printf("Get phone %s \n", iphone.GetName())

  android := androidFactory.CreatePhone("samsung")
  fmt.Printf("Get phone %s \n", android.GetName())

  other := otherFactory.CreatePhone("blackberry")
  fmt.Printf("Get phone %s \n", other.GetName())
}
