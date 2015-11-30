package bridge

import "fmt"

type MessageMobile struct {
}

func (this *MessageMobile) Send(msg string, user string) {
  fmt.Println("使用手机发送" + msg + "给" + user)
}
