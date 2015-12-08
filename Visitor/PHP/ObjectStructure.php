<?php
  class ObjectStructure {
    protected $user_list;

    public function __construct() {
      $user_list = array();
    }

    public function addUser($user) {
      $this->user_list[] = $user;
    }

    public function handleVisitor($visitor) {
      foreach($this->user_list as $user) {
          $user->accept($visitor);
      }
    }
  }
