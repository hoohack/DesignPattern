<?php
  class SecurityMgr {
    private static $instance;

    private static $user_data = array("hhq,money,look", "tom,money,get");

    private $map;

    private function __construct() {
      $this->map = array();
    }

    public static function getInstance() {
      if (self::$instance == null) {
        self::$instance = new self();
      }

      return self::$instance;
    }

    public function login($user) {
      $col = $this->queryUser($user);
      $this->map[$user] = $col;
    }

    public function hasPermit($user, $security_entity, $permit) {
      $col = $this->map[$user];
      if ($col == null || count($col) == 0) {
        echo $user . "没有登录或是没有被分配任何权限";
        return false;
      }

      foreach ($col as $fm) {
        if ($fm->match($security_entity, $permit)) {
          return true;
        }
      }

      return false;
    }

    private function queryUser($user) {
      $col = array();
      foreach (self::$user_data as $data) {
        $str = explode(",", $data);
        if ($str[0] == $user) {
            $fm = FlyweightFactory::getInstance()->getFlyweight($str[1] . ',' . $str[2]);
            $col[] = $fm;
        }
      }
      return $col;
    }
  }
