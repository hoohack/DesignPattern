package memento

type Memento struct {
  temp_record string
  temp_time int
}

func NewMemento(temp_record string, temp_time int) *Memento {
  return &Memento{temp_record, temp_time}
}

func (this *Memento) GetTempRecord() string {
  return this.temp_record
}

func (this *Memento) GetTempTime() int {
  return this.temp_time
}
