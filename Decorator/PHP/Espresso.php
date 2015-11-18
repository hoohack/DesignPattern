<?php
  class Espresso extends Beverage {
    public function __construct() {
      parent::__construct();
      $this->description = " Espresso";
    }

    public function cost() {
      return 5.6;
    }

    public function getDescription() {
      return $this->description;
    }
  }
