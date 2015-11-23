<?php
  require_once './Singleton.php';

  $instance = Singleton::getInstance();

  $another_instance = Singleton::getInstance();
  echo ($instance === $another_instance);
?>
