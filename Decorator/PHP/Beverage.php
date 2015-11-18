<?php
  abstract class Beverage {
    public $description;

    public function __construct() {
      $this->description = "Coffee";
    }

    public abstract function cost();
    public abstract function getDescription();
  }
