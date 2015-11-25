package state

import "fmt"

type BlackVoteState struct {
  votestate
}

func (this *BlackVoteState) HandleVote(user_name string, item string, vm *VoteManager) {
  fmt.Printf("你已进入黑名单")
}
