package main

import "./state"

func main() {
  vm := state.NewVoteManager()
  for i := 0; i < 9; i++ {
    vm.Vote("tom", "A")
  }
}
