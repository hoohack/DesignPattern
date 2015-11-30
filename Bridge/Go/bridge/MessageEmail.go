package bridge

import "fmt"

type MessageEmail struct {
}

func (this *MessageEmail) Send(msg string, user string) {
  fmt.Println("使用Email发送" + msg + "给" + user)
}
