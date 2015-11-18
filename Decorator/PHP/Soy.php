<?php
  class Soy extends CondimentDecorator {
    public $beverage;

    public function __construct($beverage) {
      $this->description = " Soy";
      $this->beverage = $beverage;
    }

    public function getDescription() {
      return $this->beverage->getDescription() . $this->description;
    }

    public function cost() {
      return $this->beverage->cost() + 8;
    }
  }
