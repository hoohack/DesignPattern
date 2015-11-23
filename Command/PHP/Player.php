<?php
  class Player {
    protected $state;

    public function __construct() {
      $this->state = "";
    }

    public function playing() {
      $this->state = "playing";
    }

    public function pause() {
      $this->state = "pause";
    }

    public function stop() {
      $this->state = "stop";
    }

    public function getState() {
      return $this->state;
    }
  }
