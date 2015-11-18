<?php
  class Whip extends CondimentDecorator {
    public $beverage;

    public function __construct($beverage) {
      $this->description = " Whip";
      $this->beverage = $beverage;
    }

    public function getDescription() {
      return $this->beverage->getDescription() . $this->description;
    }

    public function cost() {
      return $this->beverage->cost() + 9.9;
    }
  }
