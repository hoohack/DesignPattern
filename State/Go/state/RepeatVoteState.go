package state

import "fmt"

type RepeatVoteState struct {
  votestate
}

func (this *RepeatVoteState) HandleVote(user_name string, item string, vm *VoteManager) {
  fmt.Printf("请不要重复投票\n")
}
