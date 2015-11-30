<?php
  class CommonMessage extends AbstractMessage {
    public function __construct($impl) {
      parent::__construct($impl);
    }

    public function sendMessage($msg, $user) {
      parent::sendMessage($msg, $user);
    }
  }
