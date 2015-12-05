package memento

import "fmt"

type Workout struct {
  temp_record string
  temp_time int
}

func NewWorkout() *Workout {
  return &Workout{"", 0}
}

func (this *Workout) Running() {
  this.temp_record = "running 5 km;"
  this.temp_time = 33
}

func (this *Workout) Bicycle() {
  this.temp_record += "ride bicycle 5 km;"
  this.temp_time += 20
  this.printResult()
}

func (this *Workout) AbdominalRipperX() {
  this.temp_record += "finish abdominal ripper x;"
  this.temp_time += 30
  this.printResult()
}

func (this *Workout) CreateMemento() *Memento {
  return &Memento{this.temp_record, this.temp_time}
}

func (this *Workout) SetMemento(memento *Memento) {
  this.temp_record = memento.GetTempRecord()
  this.temp_time = memento.GetTempTime()
}

func (this *Workout) printResult() {
  fmt.Printf("%scost %d mins\n", this.temp_record, this.temp_time)
}
