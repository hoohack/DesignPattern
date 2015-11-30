package bridge

type SpecialUrgencyMessage struct {
  *AbstractMessage
}

func (this *SpecialUrgencyMessage) SendMessage(msg string, user string) {
  msg = "加急" + msg
  this.AbstractMessage.SendMessage(msg, user)
}
