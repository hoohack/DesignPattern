<?php
  class DBQueryProxy implements IDBQuery {
    private $real = null;

    public function request() {
      if ($this->real == null) {
        $this->real = new DBQuery();
      }

      return $this->real->request();
    }
  }
