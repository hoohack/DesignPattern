<?php
  class FilterChain implements Filter {
    private $filters;

    public function __construct() {
      $this->filters = array();
    }

    public function addFilter($filter) {
      array_push($this->filters, $filter);
      return $this;
    }

    public function handleFilter($str) {
      foreach ($this->filters as $filter) {
        $str = $filter->handleFilter($str);
      }

      return $str;
    }
  }
