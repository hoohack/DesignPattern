<?php
  class LifeArticle extends Article {

    public function __clone() {
      $author = new Author($this->author->getName());
      $this->setAuthor($author);
    }

    public function setAuthor($author) {
      $this->author = $author;
    }

    public function getAuthor() {
      return $this->author;
    }
  }
