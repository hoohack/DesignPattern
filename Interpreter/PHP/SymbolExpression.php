<?php
  abstract class SymbolExpression implements Expression {
    protected $left;
    protected $right;

    public function __construct($left, $right) {
      $this->left = $left;
      $this->right = $right;
    }
  }
