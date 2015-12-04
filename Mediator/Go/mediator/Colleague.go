package mediator

type Colleague struct {
  mediator Mediator
}

func NewColleague(mediator Mediator) *Colleague {
  return &Colleague{mediator}
}

func (this *Colleague) GetMediator() Mediator {
  return this.mediator
}
