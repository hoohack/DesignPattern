<?php
  class StopCommand implements Command {
    protected $player;

    public function __construct($player) {
      $this->player = $player;
    }

    public function excute() {
      $this->pre_state = $this->player->getState();
      $this->player->stop();
      echo $this->player->getState() . "\n";
    }

    public function undo() {
      if ($this->pre_state == "pause") {
        $this->player->pause();
      } else if ($this->pre_state == "playing") {
        $this->player->playing();
      } else if ($this->pre_state == "stop" ) {
        $this->player->stop();
      }

      echo $this->player->getState() . "\n";
    }
  }
