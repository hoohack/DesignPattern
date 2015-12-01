<?php
  class AuthorizationFlyweight implements Flyweight {
    private $security_entity;

    private $permit;

    public function __construct($state) {
      $arr = explode(",", $state);
      list($this->security_entity, $this->permit) = $arr;
    }

    public function getSecurityEntity() {
      return $this->security_entity;
    }

    public function getPermit() {
      return $this->permit;
    }

    public function match($security_entity, $permit) {
      if ($this->security_entity == $security_entity &&
        $this->permit == $permit) {
          return true;
      }
      return false;
    }
  }
