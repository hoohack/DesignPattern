package main

import "./bridge"

func main() {
  Impl := &bridge.MessageSMS{}
  m := &bridge.CommonMessage{&bridge.AbstractMessage{Impl}}
  m.SendMessage("hello", "tom")

  um := &bridge.UrgencyMessage{&bridge.AbstractMessage{Impl}}
  um.SendMessage("hello", "tom")

  sm := &bridge.SpecialUrgencyMessage{&bridge.AbstractMessage{Impl}}
  sm.SendMessage("hello", "tom")
}
