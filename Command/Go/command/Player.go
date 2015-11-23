package command

type Player struct {
  state string
}

func NewPlayer() *Player {
  return &Player{""}
}

func (this *Player) Playing() {
  this.state = "playing"
}

func (this *Player) Pause() {
  this.state = "pause"
}

func (this *Player) Stop() {
  this.state = "stop"
}

func (this *Player) GetState() string {
  return this.state
}
