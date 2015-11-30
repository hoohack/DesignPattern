<?php
  class MessageEmail implements MessageImplementor {
    public function send($msg, $user) {
      echo "使用Email的方式，发送信息" . $msg . "给" . $user . "\n";
    }
  }
