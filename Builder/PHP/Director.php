<?php
  class Director {
    private $builder;

    public function __construct($builder) {
      $this->builder = $builder;
    }

    public function construct($header, $body, $footer) {
      $this->builder->buildHeader($header);
      $this->builder->buildBody($body);
      $this->builder->buildFooter($footer);
    }
  }
