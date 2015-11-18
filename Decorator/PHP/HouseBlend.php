<?php
  class HouseBlend extends Beverage {
    public function __construct() {
      parent::__construct();
      $this->description = " HouseBlend";
    }

    public function cost() {
      return 8.9;
    }

    public function getDescription() {
      return $this->description;
    }
  }
