<?php
  abstract class Article {
    protected $title;

    protected $author;

    abstract public function __clone();

    public function getTitle() {
      return $this->title;
    }

    public function setTitle($title) {
      $this->title = $title;
    }
  }
