<?php
  class BlackVoteState implements VoteState {
    public function handleVote($user, $vote_item, $vote_manager) {
      echo "你已进入黑名单\n";
    }
  }
