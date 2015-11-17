<?php
  class ObserverB implements Observer
  {
    public function update(Subject $subject)
    {
      echo "observer B value is " . $subject->getValue() . "\n";
    }
  }
?>
