<?php
  interface MessageImplementor {
    public function send($msg, $user);
  }
