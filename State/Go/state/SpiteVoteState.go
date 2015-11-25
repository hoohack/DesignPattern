package state

import "fmt"

type SpiteVoteState struct {
  votestate
}

func (this *SpiteVoteState) HandleVote(user_name string, item string, vm *VoteManager) {
  if _, ok := vm.Votes[user_name];ok {
    delete(vm.Votes, user_name)
  }

  fmt.Printf("你有恶意刷票行为，取消投票资格\n")
}
