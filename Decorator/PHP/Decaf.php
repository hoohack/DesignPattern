<?php
  class Decaf extends Beverage {
    public function __construct() {
      parent::__construct();
      $this->description = " Decaf";
    }

    public function cost() {
      return 6.0;
    }

    public function getDescription() {
      return $this->description;
    }
  }
