<?php
  class View extends Colleague {
    public function makeAddRequest($user) {
      echo "发起增加用户{$user}请求\n";
      $this->getMeditor()->addUser($user);
    }

    public function showUser() {
      print_r($this->getMeditor()->getUserList());
    }

    public function makeDeleteRequest($user) {
      echo "发起删除用户{$user}请求\n";
      $this->getMeditor()->deleteUser($user);
    }

    public function showResult($result) {
      echo $result;
    }
  }
