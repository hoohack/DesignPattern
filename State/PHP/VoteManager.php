<?php
  class VoteManager {
    public $votes;
    private $vote_count;
    private $state;

    public function __construct() {
      $this->votes = array();
      $this->vote_count = array();
    }

    public function add($user, $item) {
      $this->vote[] = array("user" => $user, "item" => $item);
    }

    public function vote($user, $item) {
      if (!isset($this->vote_count[$user])) {
        $this->vote_count[$user] = 0;
      }
      $current_count = ++$this->vote_count[$user];

      if ($current_count == 1) {
        $this->state = new NormalVoteState();
      } elseif($current_count > 1 && $current_count < 5) {
        $this->state = new RepeatVoteState();
      } elseif ($current_count >= 5 && $current_count < 8) {
        $this->state = new SpiteVoteState();
      } elseif ($current_count > 8) {
        $this->state = new BlackVoteState();
      }

      $this->state->handleVote($user, $item, $this);
    }
  }
