<?php
  class FlyweightFactory {
    private static $instance;

    private $fwMap;

    private function __construct() {
      $this->fwMap = array();
    }

    public static function getInstance() {
      if (self::$instance === null) {
        self::$instance = new self();
      }

      return self::$instance;
    }

    public function getFlyweight($key) {
      if (!isset(self::$instance->fwMap[$key])) {
        self::$instance->fwMap[$key] = new AuthorizationFlyweight($key);
      }

      return self::$instance->fwMap[$key];
    }

  }
