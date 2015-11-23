package command

type PlayCommand struct {
  command
  *Player
  pre_state string
}

func NewPlayCommand(p *Player) *PlayCommand {
  return &PlayCommand{Player: p, pre_state: ""}
}

func (this *PlayCommand) Excute() string {
  this.pre_state = this.Player.GetState()
  this.Player.Playing()
  return this.Player.GetState()
}

func (this *PlayCommand) Undo() string {
  if this.pre_state == "pause" {
    this.Player.Pause()
  } else if this.pre_state == "playing" {
    this.Player.Playing()
  } else if this.pre_state == "stop" {
    this.Player.Stop()
  }

  return this.Player.GetState()
}
