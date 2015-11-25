package state

import "fmt"

type NormalVoteState struct {
  votestate
}

func (this *NormalVoteState) HandleVote(user_name string, item string, vm *VoteManager) {
  vm.Add(user_name, item)
  fmt.Printf("投票成功\n")
}
