<?php
  interface VoteState {
    public function handleVote($user, $vote_item, $vote_manager);
  }
