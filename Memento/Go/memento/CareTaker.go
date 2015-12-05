package memento

type CareTaker struct {
  memento *Memento
}

func (this *CareTaker) SaveMemento(memento *Memento) {
  this.memento = memento
}

func (this *CareTaker) RetriveMemento() *Memento {
  return this.memento
}
