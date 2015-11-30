<?php
  class MessageMobile implements MessageImplementor {
    public function send($msg, $user) {
      echo "使用手机，发送信息" . $msg . "给" . $user . "\n";
    }
  }
