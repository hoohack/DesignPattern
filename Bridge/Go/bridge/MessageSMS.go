package bridge

import "fmt"

type MessageSMS struct {
}

func (this *MessageSMS) Send(msg string, user string) {
  fmt.Println("使用站内短消息发送" + msg + "给" + user)
}
