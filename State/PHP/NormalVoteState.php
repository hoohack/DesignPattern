<?php
  class NormalVoteState implements VoteState {
    public function handleVote($user, $vote_item, $vote_manager) {
      $vote_manager->add($user, $vote_item);
      echo "投票成功\n";
    }
  }
