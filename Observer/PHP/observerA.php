<?php
  class ObserverA implements Observer
  {
    public function update(Subject $subject)
    {
      echo "observer A value is " . $subject->getValue() . "\n";
    }
  }
?>
