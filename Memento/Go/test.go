package main

import "./memento"

func main() {
  workout := memento.NewWorkout()
  workout.Running()
  care_taker := &memento.CareTaker{}
  memento := workout.CreateMemento()
  care_taker.SaveMemento(memento)

  workout.Bicycle()
  workout.SetMemento(care_taker.RetriveMemento())
  workout.AbdominalRipperX()
}
