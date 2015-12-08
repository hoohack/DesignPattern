<?php
  abstract class User {
    private $name;
    private $point;

    public function __construct($name) {
      $this->name = $name;
      $this->point = 0;
    }

    public function addPoint($point) {
      $this->point += $point;
    }

    public function getName() {
      return $this->name;
    }

    public function getPoint() {
      return $this->point;
    }

    abstract function accept($visitor);
  }
