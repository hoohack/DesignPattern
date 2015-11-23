package command

type Invoker struct {
  command command
}

func (this *Invoker) SetCommand(cmd command) {
  this.command = cmd
}

func (this *Invoker) Run() string {
  return this.command.Excute()
}

func (this *Invoker) Undo() string {
  return this.command.Undo()
}
