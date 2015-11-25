<?php
  class RepeatVoteState implements VoteState {
    public function handleVote($user, $vote_item, $vote_manager) {
      echo "请不要重复投票\n";
    }
  }
