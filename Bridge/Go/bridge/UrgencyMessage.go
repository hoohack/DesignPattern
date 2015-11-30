package bridge

type UrgencyMessage struct {
  *AbstractMessage
}

func (this *UrgencyMessage) SendMessage(msg string, user string) {
  msg = "紧急" + msg
  this.AbstractMessage.SendMessage(msg, user)
}
