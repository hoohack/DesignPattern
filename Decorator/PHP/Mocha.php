<?php
  class Mocha extends CondimentDecorator {
    public $beverage;

    public function __construct($beverage) {
      $this->description = ' Mocha';
      $this->beverage = $beverage;
    }

    public function cost() {
      return $this->beverage->cost() + 10.3;
    }

    public function getDescription() {
      return $this->beverage->getDescription() . $this->description;
    }
  }
