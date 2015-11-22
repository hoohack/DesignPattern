<?php
  class SpecialFactory {
    protected $name;

    public function __construct($name) {
      $this->name = $name;
    }

    public function getName() {
      return $this->name;
    }
  }
?>
