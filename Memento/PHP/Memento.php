<?php
  class Memento {
    private $temp_record;

    private $temp_time;

    public function __construct($temp_record, $temp_time) {
      $this->temp_record = $temp_record;
      $this->temp_time = $temp_time;
    }

    public function getTempRecord() {
      return $this->temp_record;
    }

    public function getTempTime() {
      return $this->temp_time;
    }
  }
