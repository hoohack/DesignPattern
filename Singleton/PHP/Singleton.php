<?php
  class Singleton {
    private static $_instance;

    private function __construct() {}

    public static function getInstance() {
      if (self::$_instance === NULL) {
        self::$_instance = new Singleton();
      }

      return self::$_instance;
    }
  }
?>
