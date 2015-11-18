<?php
  class DarkRoast extends Beverage {
    public function __construct() {
      parent::__construct();
      $this->description = " DarkRoast";
    }

    public function cost() {
      return 5.0;
    }

    public function getDescription() {
      return $this->description;
    }
  }
