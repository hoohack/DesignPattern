<?php
  abstract class Special implements PhoneFactory {
    protected $content;

    public function __construct($content) {
      $this->content = $content;
    }
  }
?>
