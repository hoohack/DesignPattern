package state

type votestate interface {
  HandleVote(user_name string, item string, vote_manager *VoteManager);
}
