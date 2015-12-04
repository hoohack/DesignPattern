<?php
  class Model extends Colleague {
    private $user_list;

    public function __construct($mediator) {
      parent::__construct($mediator);
      $user_list = array();
    }

    public function addUser($user) {
      $this->user_list[] = $user;
    }

    public function getUserList() {
      return $this->user_list;
    }

    public function deleteUser($user) {
      foreach ($this->user_list as $key => $user_name) {
        if ($user_name == $user) {
          unset($this->user_list[$key]);
          $this->getMeditor()->showResult("删除" . $user . "成功\n");
          return;
        }
      }

      $this->getMeditor()->showResult("删除" . $user . "失败，该用户不存在\n");
    }
  }
