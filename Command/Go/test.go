package main

import "fmt"
import "./command"

func main() {
  invoker := new(command.Invoker)
  player := command.NewPlayer()

  invoker.SetCommand(command.NewPlayCommand(player))
  fmt.Printf("%s\n", invoker.Run())

  invoker.SetCommand(command.NewPauseCommand(player))
  fmt.Printf("%s\n", invoker.Run())

  fmt.Printf("%s\n", invoker.Undo())

  invoker.SetCommand(command.NewStopCommand(player))
  fmt.Printf("%s\n", invoker.Run())
}
