<?php
  abstract class AbstractMessage {
    protected $impl;

    public function __construct($impl) {
      $this->impl = $impl;
    }

    public function sendMessage($msg, $user) {
      $this->impl->send($msg, $user);
    }
  }
