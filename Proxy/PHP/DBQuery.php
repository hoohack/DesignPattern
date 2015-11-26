<?php
  class DBQuery implements IDBQuery {
    public function __construct() {
      echo "connecting...\n";
    }

    public function request() {
      echo "requesting...\n";
    }
  }
