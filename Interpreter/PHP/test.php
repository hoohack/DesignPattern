<?php
      require_once './Expression.php';
      require_once './ValueExpression.php';
      require_once './SymbolExpression.php';
      require_once './PlusExpression.php';
      require_once './MinusExpression.php';
      require_once './MulExpression.php';
      require_once './DivExpression.php';
      require_once './ModExpression.php';
      require_once './Calculator.php';


      $statement = "3 * 2 * 4";

      $calculator = new Calculator();

      $calculator->init($statement);

      $result = $calculator->compute();

      echo $statement . " = " . $result . "\n";
