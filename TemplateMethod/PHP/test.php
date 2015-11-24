<?php
  require_once './ViewEngine.php';
  require_once './JsonView.php';
  require_once './HtmlView.php';

  $json_view = new JsonView();
  $json_view->render();

  $html_view = new HtmlView();
  $html_view->render();
