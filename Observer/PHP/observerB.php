<?php
  class ObserverB implements Observer
  {
    private $_value;

    public function update($value)
    {
      $this->_value = $value;
      $this->display();
    }

    public function display() {
      echo "Observer B value is " . $this->_value . "\n";
    }
  }
?>
