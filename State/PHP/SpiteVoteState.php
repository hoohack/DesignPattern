<?php
  class SpiteVoteState implements VoteState {
    public function handleVote($user, $vote_item, $vote_manager) {
      foreach ($vote_manager->votes as $key => $vote) {
        if ($vote_manager->vote["user"] == $user) {
          unset($vote_manager->vote[$key]);
          break;
        }
      }

      echo "你有恶意刷票行为，取消投票资格\n";
    }
  }
