package bridge

type AbstractMessage struct {
  Impl MessageImplementor
}

func (this *AbstractMessage) SendMessage(msg string, user string) {
  this.Impl.Send(msg, user)
}
