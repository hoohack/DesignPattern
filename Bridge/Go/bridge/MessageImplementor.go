package bridge

type MessageImplementor interface {
  Send(msg string, user string)
}
