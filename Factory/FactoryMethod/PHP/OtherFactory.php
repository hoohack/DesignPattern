<?php
  class OtherFactory extends PhoneFactory {
    public function createPhone($type) {
      if ($type == "nokia") {
        return new Nokia();
      } else if ($type == "blackberry") {
        return new BlackBerry();
      }
    }
  }
?>
