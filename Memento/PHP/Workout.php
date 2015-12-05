<?php
  class Workout {

    private $temp_record;

    private $temp_time;

    public function __construct() {
      $this->temp_record = "";
      $this->temp_time = 0;
    }

    public function running() {
      $this->temp_record = "running 5 km;";
      $this->temp_time = 33;
    }

    public function bicycle() {
      $this->temp_record .= "ride bicycle 5 km;";
      $this->temp_time += 20;
      // echo $this->temp_record . " cost " . $this->temp_time . "mins\n";
    }

    public function abdominal_ripper_x() {
      $this->temp_record .= "finish abdominal ripper x;";
      $this->temp_time += 30;
      echo $this->temp_record . " cost " . $this->temp_time . "mins\n";
    }

    public function createMemento() {
      return new Memento($this->temp_record, $this->temp_time);
    }

    public function setMemento($memento) {
      $this->temp_record = $memento->getTempRecord();
      $this->temp_time = $memento->getTempTime();
    }
  }
