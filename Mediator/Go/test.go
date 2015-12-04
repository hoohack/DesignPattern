package main

import "./mediator"

func main() {
  controller := &mediator.Controller{}
  view := mediator.NewView(controller)
  model := mediator.NewModel(controller)
  controller.SetColleague(model, view)
  view.MakeAddRequest("tom")
  view.MakeAddRequest("amy")
  view.ShowUser()
  view.MakeDeleteRequest("tom")
  view.ShowUser()
  view.MakeDeleteRequest("aaa")
}
