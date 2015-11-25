<?php
  require_once './VoteState.php';
  require_once './NormalVoteState.php';
  require_once './RepeatVoteState.php';
  require_once './SpiteVoteState.php';
  require_once './BlackVoteState.php';
  require_once './VoteManager.php';

  $vote_manager = new VoteManager();
  for($i=0;$i<9;$i++){
      $vote_manager->vote("u1","A");
  }
