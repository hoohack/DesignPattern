<?php
  class Controller implements Mediator {
    protected $model;

    protected $view;

    public function setColleague($model, $view) {
      $this->model = $model;
      $this->view = $view;
    }

    public function addUser($user) {
      $this->model->addUser($user);
    }

    public function getUserList() {
      return $this->model->getUserList();
    }

    public function deleteUser($user) {
      $this->model->deleteUser($user);
    }

    public function showResult($result) {
      $this->view->showResult($result);
    }
  }
