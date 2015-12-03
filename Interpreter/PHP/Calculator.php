<?php
  class Calculator {
    private $statement;
    private $expression;
    private $symbol_arr = array("+", "-", "*", "/", "%");
    private $symbol_map = array("+" => "PlusExpression", "-" => "MinusExpression",
      "*" => "MulExpression", "/" => "DivExpression", "%" => "ModExpression");

    public function init($statement) {
      $expression_stack = array();
      $left = null;
      $right = null;

      $statement_arr = explode(" ", $statement);
      for ($i = 0; $i < count($statement_arr); ++$i) {
        echo "run";
        $elem = $statement_arr[$i];
        if (in_array($elem, $this->symbol_arr)) {
          $left = array_pop($expression_stack);
          $right = new ValueExpression($statement_arr[++$i]);
          print_r($left);
          print_r($right);
          array_push($expression_stack, new $this->symbol_map[$elem]($left, $right));
        } else {
          array_push($expression_stack, new ValueExpression($elem));
        }
      }
      $this->expression = array_pop($expression_stack);
    }

    public function compute() {
      return $this->expression->calculate();
    }
  }
