<?php
  class CareTaker {
    private $memento;

    public function saveMemento($memento) {
      $this->memento = $memento;
    }

    public function retriveMemento() {
      return $this->memento;
    }
  }
