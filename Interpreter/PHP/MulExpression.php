<?php
  class MulExpression extends SymbolExpression {
    public function __construct($left, $right) {
      parent::__construct($left, $right);
    }

    public function calculate() {
      return $this->left->calculate() * $this->right->calculate();
    }
  }
