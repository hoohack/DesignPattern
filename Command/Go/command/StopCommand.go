package command

type StopCommand struct {
  command
  *Player
  pre_state string
}

func NewStopCommand(p *Player) *StopCommand {
  return &StopCommand{Player: p, pre_state: ""}
}

func (this *StopCommand) Excute() string {
  this.pre_state = this.Player.GetState()
  this.Player.Stop()
  return this.Player.GetState()
}

func (this *StopCommand) Undo() string {
  if this.pre_state == "pause" {
    this.Player.Pause()
  } else if this.pre_state == "playing" {
    this.Player.Playing()
  } else if this.pre_state == "stop" {
    this.Player.Stop()
  }

  return this.Player.GetState()
}
