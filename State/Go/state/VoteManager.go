package state

type VoteManager struct {
  Votes map[string]string
  vote_count map[string]int
  state votestate
}

func NewVoteManager() *VoteManager {
  return &VoteManager{Votes: make(map[string]string), vote_count: make(map[string]int)}
}

func (this *VoteManager) Add(user_name string, item string) {
  this.Votes[user_name] = item
}

func (this *VoteManager) Vote(user_name string, item string) {
  if _, ok := this.vote_count[user_name];!ok {
    this.vote_count[user_name] = 0
  }

  this.vote_count[user_name]++
  current_count := this.vote_count[user_name]
  if current_count == 1 {
    this.state = &NormalVoteState{}
  } else if current_count > 1 && current_count < 5 {
    this.state = &RepeatVoteState{}
  } else if current_count > 5 && current_count < 8 {
    this.state = &SpiteVoteState{}
  } else if current_count > 8 {
    this.state = &BlackVoteState{}
  }

  this.state.HandleVote(user_name, item, this)
}
