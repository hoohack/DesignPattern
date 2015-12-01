<?php
  class ValueExpression implements Expression {
    private $value;

    public function __construct($value) {
      $this->value = $value;
    }

    public function calculate() {
      return $this->value;
    }
  }
