<?php
  class Person
  {
    private $_strategy = null;

    public function setStrategy(TravelStrategy $strategy)
    {
      $this->_strategy = $strategy;
    }

    public function travel()
    {
      $this->_strategy->travelAlgorithm();
    }
  }
?>
