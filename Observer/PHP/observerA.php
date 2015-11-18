<?php
  class ObserverA implements Observer
  {
    private $_value;

    public function update($value)
    {
      $this->_value = $value;
      $this->display();
    }

    public function display() {
      echo "Observer A value is " . $this->_value . "\n";
    }
  }
?>
