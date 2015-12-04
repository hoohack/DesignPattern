<?php
  interface Mediator {
    public function addUser($user);

    public function getUserList();

    public function deleteUser($user);

    public function showResult($result);
  }
