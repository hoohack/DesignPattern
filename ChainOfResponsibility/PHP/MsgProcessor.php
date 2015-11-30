<?php
  class MsgProcessor {
    private $msg;
    private $filter_chain;

    public function __construct($msg, $filter_chain) {
      $this->msg = $msg;
      $this->filter_chain = $filter_chain;
    }

    public function process() {
      return $this->filter_chain->handleFilter($this->msg);
    }
  }
