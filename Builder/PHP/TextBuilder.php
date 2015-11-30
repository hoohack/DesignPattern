<?php
  class TextBuilder implements Builder {
    private $content;

    public function buildHeader($header) {
      $this->content .= "Text header: " . $header . "\n";
    }

    public function buildBody($body) {
      $this->content .= "Text body: " . $body . "\n";
    }

    public function buildFooter($footer) {
      $this->content .= "Text footer: " . $footer . "\n";
    }

    public function getResult() {
      return $this->content;
    }
  }
