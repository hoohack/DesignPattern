package main

import (
  "fmt"
  "./factorymethod"
)

func main() {
  iphoneFactory := &factorymethod.IPhoneFactory{}
  androidFactory := &factorymethod.AndroidFactory{}
  otherFactory := &factorymethod.OtherFactory{}

  iphone := iphoneFactory.CreatePhone()
  fmt.Printf("Get phone %s \n", iphone.GetName())

  android := androidFactory.CreatePhone()
  fmt.Printf("Get phone %s \n", android.GetName())

  other := otherFactory.CreatePhone()
  fmt.Printf("Get phone %s \n", other.GetName())
}
