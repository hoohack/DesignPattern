package main

import "fmt"
import "./abstractfactory"

func main() {
  iPhoneFactory := new(abstractfactory.IPhoneFactory)
  iphone_os := iPhoneFactory.CreateOS("IOS")
  fmt.Printf("Creating %s...\n", iphone_os.Create())
  iphone_special := iPhoneFactory.CreateSpecial("siri")
  fmt.Printf("Creating %s...\n", iphone_special.Create())

  androidFactory := new(abstractfactory.AndroidFactory)
  android_os := androidFactory.CreateOS("Android")
  fmt.Printf("Creating %s...\n", android_os.Create())
  android_special := androidFactory.CreateSpecial("nfc")
  fmt.Printf("Creating %s...\n", android_special.Create())
}
