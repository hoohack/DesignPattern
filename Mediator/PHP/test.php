<?php
  require_once './Colleague.php';
  require_once './Mediator.php';
  require_once './Controller.php';
  require_once './Model.php';
  require_once './View.php';

  $controller = new Controller();
  $view = new View($controller);
  $model = new Model($controller);
  $controller->setColleague($model, $view);
  $view->makeAddRequest('tom');
  $view->makeAddRequest('amy');
  $view->showUser();
  $view->makeDeleteRequest("tom");
  $view->showUser();
  $view->makeDeleteRequest("aaa");
