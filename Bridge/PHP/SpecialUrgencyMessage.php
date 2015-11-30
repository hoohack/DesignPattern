<?php
  class SpecialUrgencyMessage extends AbstractMessage {
    public function __construct($impl) {
      parent::__construct($impl);
    }

    public function sendMessage($msg, $user) {
      $msg = "特急信息:" . $msg;
      parent::sendMessage($msg, $user);
    }
  }
