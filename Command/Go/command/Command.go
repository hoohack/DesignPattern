package command

type command interface {
  Excute() string
  Undo() string
}
