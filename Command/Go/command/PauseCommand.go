package command

type PauseCommand struct {
  command
  *Player
  pre_state string
}

func NewPauseCommand(p *Player) *PauseCommand {
  return &PauseCommand{Player: p, pre_state: ""}
}

func (this *PauseCommand) Excute() string {
  this.pre_state = this.Player.GetState()
  this.Player.Pause()
  return this.Player.GetState()
}

func (this *PauseCommand) Undo() string {
  if this.pre_state == "pause" {
    this.Player.Pause()
  } else if this.pre_state == "playing" {
    this.Player.Playing()
  } else if this.pre_state == "stop" {
    this.Player.Stop()
  }

  return this.Player.GetState()
}
