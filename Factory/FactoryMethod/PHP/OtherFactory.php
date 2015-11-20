<?php
  class OtherFactory implements PhoneFactory {
    public function createPhone() {
      return new Other();
    }
  }
?>
