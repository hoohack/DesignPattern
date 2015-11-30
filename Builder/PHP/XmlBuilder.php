<?php
  class XmlBuilder implements Builder {
    private $content;

    public function buildHeader($header) {
      $this->content .= "<Header>" . $header . "</Header>\n";
    }

    public function buildBody($body) {
      $this->content .= "<Body>" . $body . "</Body>\n";
    }

    public function buildFooter($footer) {
      $this->content .= "<Footer>" . $footer . "</Footer>\n";
    }

    public function getResult() {
      return $this->content;
    }
  }
