<?php
  abstract class Colleague {
    private $mediator;

    public function __construct($mediator) {
      $this->mediator = $mediator;
    }

    protected function getMeditor() {
      return $this->mediator;
    }
  }
