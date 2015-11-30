<?php
  class MessageSMS implements MessageImplementor {
    public function send($msg, $user) {
      echo "使用站内短信息的方式，发送信息" . $msg . "给" . $user . "\n";
    }
  }
