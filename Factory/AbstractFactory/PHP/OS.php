<?php
  abstract class OS implements PhoneFactory {
    protected $name;

    public function __construct($name) {
      $this->name = $name;
    }
  }
?>
