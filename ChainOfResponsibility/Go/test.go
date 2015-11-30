package main

import "./chain"
import "fmt"

func main() {
  msg := "<p>hello, i am tom.how are you?</p>"
  filter_chain := chain.NewFilterChain()
  filter_chain.AddFilter(&chain.HTMLFilter{}).AddFilter(&chain.LengthFilter{})
  msg_processor := chain.NewMsgProcessor(msg, filter_chain)
  fmt.Println(msg_processor.Process())
}
