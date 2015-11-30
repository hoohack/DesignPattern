<?php
  class UrgencyMessage extends AbstractMessage {
    public function __construct($impl) {
      parent::__construct($impl);
    }

    public function sendMessage($msg, $user) {
      $msg = "紧急信息：" . $msg;
      parent::sendMessage($msg, $user);
    }
  }
