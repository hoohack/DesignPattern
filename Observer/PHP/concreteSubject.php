<?php
  class ConcreteSubject implements Subject
  {
    private $_observers, $_value;

    public function __construct()
    {
      $this->_observers = array();
    }

    public function registerObserver(Observer $observer)
    {
      array_push($this->_observers, $observer);
    }

    public function removeObserver(Observer $observer)
    {
      if ($index = array_search($observer, $this->_observers, true))
      {
        unset($this->_observers[$index]);
      }
    }

    public function notify()
    {
      foreach ($this->_observers as $observer)
      {
        $observer->update($this->_value);
      }
    }

    public function setValue($value)
    {
      $this->_value = $value;
      $this->notify();
    }

    public function getValue()
    {
      return $this->_value;
    }
  }
?>
