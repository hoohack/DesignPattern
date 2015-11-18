<?php
  class Mocha extends CondimentDecorator {
    public $_beverage;

    public function __construct($beverage) {
      $this->description = ' Mocha';
      $this->_beverage = $beverage;
    }

    public function cost() {
      return $this->_beverage->cost() + 10.3;
    }

    public function getDescription() {
      return $this->_beverage->getDescription() . $this->description;
    }
  }
