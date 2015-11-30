package bridge

type CommonMessage struct {
  *AbstractMessage
}

func (this *CommonMessage) SendMessage(msg string, user string) {
  this.AbstractMessage.SendMessage(msg, user)
}
